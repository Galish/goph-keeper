package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (s *psqlStore) UpdateSecureNote(ctx context.Context, note *repository.SecureNote) error {
	protected, err := s.encrypt(note)
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
		note.Title,
		note.Description,
		protected,
		note.LastEditedAt,
		note.ID,
		note.Type,
		note.CreatedBy,
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

	if note.Version > 0 && note.Version != version {
		tx.Rollback()

		return repository.ErrVersionConflict
	}

	return tx.Commit()
}
