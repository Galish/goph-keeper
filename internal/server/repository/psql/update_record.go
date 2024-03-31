package psql

import (
	"context"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (s *psqlStore) UpdateSecureRecord(
	ctx context.Context,
	record *repository.SecureRecord,
) error {
	res, err := s.db.ExecContext(
		ctx,
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
				last_edited_at = $14
			WHERE uuid = $1
				AND type = $2
				AND created_by = $13
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
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count < 1 {
		return repository.ErrNotFound
	}

	return err
}
