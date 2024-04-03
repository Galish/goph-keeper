package psql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (s *psqlStore) UpdateSecureRecord(ctx context.Context, record *repository.SecureRecord) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	row := tx.QueryRow(
		`
			UPDATE secure_notes
			SET title = $3,
				description = $4,
				username = $5,
				password = $6,
				text_note = $7,
				raw_note = $8,
				card_number = $9,
				card_holder = $10,
				card_cvc = $11,
				card_expiry = $12,
				last_edited_at = $14,
				version = version + 1
			WHERE uuid = $1
				AND type = $2
				AND created_by = $13
			RETURNING version
		`,
		record.ID,
		record.Type,
		record.Title,
		record.Description,
		record.Username,
		record.Password,
		record.TextNote,
		record.RawNote,
		record.CardNumber,
		record.CardHolder,
		record.CardCVC,
		record.CardExpiry,
		record.CreatedBy,
		record.LastEditedAt,
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
