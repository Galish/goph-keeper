package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (s *Store) GetSecureNote(
	ctx context.Context,
	user, id string,
	noteType repository.SecureNoteType,
) (*repository.SecureNote, error) {
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
		noteType,
	)

	var (
		note      repository.SecureNote
		protected string
	)

	err := row.Scan(
		&note.ID,
		&note.Type,
		&note.Title,
		&note.Description,
		&protected,
		&note.CreatedBy,
		&note.CreatedAt,
		&note.LastEditedAt,
		&note.Version,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	if err := s.decrypt(protected, &note); err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}

	return &note, nil
}
