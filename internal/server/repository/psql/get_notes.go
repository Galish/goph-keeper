package psql

import (
	"context"
	"fmt"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (s *psqlStore) GetSecureNotes(ctx context.Context, user string, noteType repository.SecureNoteType) ([]*repository.SecureNote, error) {
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
		noteType, user,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var (
		notes     []*repository.SecureNote
		protected string
	)

	for rows.Next() {
		var note repository.SecureNote

		if err := rows.Scan(
			&note.ID,
			&note.Type,
			&note.Title,
			&note.Description,
			&protected,
			&note.CreatedBy,
			&note.CreatedAt,
			&note.LastEditedAt,
			&note.Version,
		); err != nil {
			return nil, err
		}

		if err := s.decrypt(protected, &note); err != nil {
			return nil, fmt.Errorf("decryption failed: %v", err)
		}

		notes = append(notes, &note)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}
