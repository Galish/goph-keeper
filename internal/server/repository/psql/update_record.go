package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (s *psqlStore) UpdateSecureRecord(ctx context.Context, record *repository.SecureRecord) error {
	protected, err := s.encrypt(record)
	if err != nil {
		return fmt.Errorf("encryption failed: %v", err)
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	row := tx.QueryRow(
		`
			UPDATE secure_notes
			SET
				title = $1,
				description = $2,
				protected_data = $3,
				last_edited_at = $4,
				version = version + 1
			WHERE
				uuid = $5
				AND type = $6
				AND created_by = $7
			RETURNING version
		`,
		record.Title,
		record.Description,
		protected,
		record.LastEditedAt,
		record.ID,
		record.Type,
		record.CreatedBy,
	)

	var version int32
	err = row.Scan(&version)
	if errors.Is(err, sql.ErrNoRows) {
		return repository.ErrNotFound
	}

	if err != nil {
		tx.Rollback()

		return err
	}

	if record.Version > 0 && record.Version != version {
		tx.Rollback()

		return repository.ErrVersionConflict
	}

	return tx.Commit()
}
