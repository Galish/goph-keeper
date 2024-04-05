package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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
				protected_data,
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

	var (
		record    repository.SecureRecord
		protected string
	)

	err := row.Scan(
		&record.ID,
		&record.Type,
		&record.Title,
		&record.Description,
		&protected,
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

	if err := s.decrypt(protected, &record); err != nil {
		return nil, fmt.Errorf("decryption failed: %v", err)
	}

	return &record, nil
}
