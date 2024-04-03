package psql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (s *psqlStore) GetSecureRecord(
	ctx context.Context,
	user, id string,
	recordType repository.SecureRecordType,
) (*repository.SecureRecord, error) {
	row := s.db.QueryRowContext(
		ctx,
		`
			SELECT
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
				last_edited_at,
				version
			FROM
				secure_notes
			WHERE
				uuid = $1
				AND
				created_by = $2
				AND
				type = $3
			;
		`,
		id,
		user,
		recordType,
	)

	var record repository.SecureRecord
	err := row.Scan(
		&record.ID,
		&record.Type,
		&record.Title,
		&record.Description,
		&record.Username,
		&record.Password,
		&record.TextNote,
		&record.RawNote,
		&record.CardNumber,
		&record.CardHolder,
		&record.CardCVC,
		&record.CardExpiry,
		&record.CreatedBy,
		&record.CreatedAt,
		&record.LastEditedAt,
		&record.Version,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return &record, nil
}
