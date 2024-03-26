package psql

import (
	"context"
	"errors"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"

	"github.com/jackc/pgx/v5/pgconn"
)

func (s *psqlStore) SetUser(ctx context.Context, user *entity.User) error {
	_, err := s.db.ExecContext(
		ctx,
		`
			INSERT INTO users (uuid, login, password, is_active)
			VALUES ($1, $2, $3, $4);
		`,
		user.ID,
		user.Login,
		user.Password,
		user.IsActive,
	)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == errCodeConflict {
		return repository.ErrUserConflict
	}
	if err != nil {
		return err
	}

	return nil
}
