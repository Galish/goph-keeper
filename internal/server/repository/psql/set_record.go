package psql

import (
	"context"
	"errors"

	"github.com/Galish/goph-keeper/internal/server/repository"
	"github.com/jackc/pgx/v5/pgconn"
)

func (s *psqlStore) AddSecureRecord(
	ctx context.Context,
	record *repository.SecureRecord,
) error {
	_, err := s.db.ExecContext(
		ctx,
		`
			INSERT INTO secure_notes (
				uuid,
				type,
				title,
				description,
				username,
				password,
				text_note,
				raw_note,
				card_number,
				card_holder,
				card_cvc,
				card_expiry,
				created_by,
				created_at,
				last_edited_at
			)
			VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
			)
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
		record.CreatedAt,
		record.LastEditedAt,
	)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == errCodeConflict {
		return repository.ErrRecordConflict
	}
	if err != nil {
		return err
	}

	return nil
}
