package psql

import (
	"context"
	"fmt"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (s *psqlStore) GetSecureRecords(ctx context.Context, user string, recordType repository.SecureRecordType) ([]*repository.SecureRecord, error) {
	rows, err := s.db.QueryContext(
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
				type = $1
				AND
				created_by = $2
			;
		`,
		recordType, user,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		records   []*repository.SecureRecord
		protected string
	)

	for rows.Next() {
		var record repository.SecureRecord

		if err := rows.Scan(
			&record.ID,
			&record.Type,
			&record.Title,
			&record.Description,
			&protected,
			&record.CreatedBy,
			&record.CreatedAt,
			&record.LastEditedAt,
			&record.Version,
		); err != nil {
			return nil, err
		}

		if err := s.decrypt(protected, &record); err != nil {
			return nil, fmt.Errorf("decryption failed: %v", err)
		}

		records = append(records, &record)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
