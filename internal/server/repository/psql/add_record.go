package psql

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (s *psqlStore) AddSecureRecord(ctx context.Context, record *repository.SecureRecord) error {
	protected, err := s.encrypt(record)
	if err != nil {
		return fmt.Errorf("encryption failed: %v", err)
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
		record.ID,
		record.Type,
		record.Title,
		record.Description,
		protected,
		record.CreatedBy,
		record.CreatedAt,
		record.LastEditedAt,
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
