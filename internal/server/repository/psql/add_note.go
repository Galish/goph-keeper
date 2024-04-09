package psql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (s *Store) AddSecureNote(ctx context.Context, note *repository.SecureNote) error {
	protected, err := s.encrypt(note)
	if err != nil {
		return fmt.Errorf("encryption failed: %w", err)
	}

	_, err = s.db.ExecContext(
		ctx,
		`
			INSERT INTO secure_notes (
				uuid,
				type,
				title,
				description,
				protected_data,
				created_by,
				created_at,
				last_edited_at
			)
			VALUES (
				$1, $2, $3, $4, $5, $6, $7, $8
			)
		`,
		note.ID,
		note.Type,
		note.Title,
		note.Description,
		protected,
		note.CreatedBy,
		note.CreatedAt,
		note.LastEditedAt,
	)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == errCodeConflict {
		return repository.ErrConflict
	}

	if err != nil {
		return err
	}

	return nil
}
