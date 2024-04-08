package psql

import (
	"context"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (s *psqlStore) DeleteSecureNote(ctx context.Context, user, id string, noteType repository.SecureNoteType) error {
	res, err := s.db.ExecContext(
		ctx,
		`
			DELETE
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
		id, user, noteType,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return repository.ErrNotFound
	}

	return nil
}
