package psql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (s *psqlStore) GetUserByLogin(ctx context.Context, login string) (*entity.User, error) {
	row := s.db.QueryRowContext(
		ctx,
		`SELECT uuid, login, password, is_active FROM users WHERE login=$1;`,
		login,
	)

	var user entity.User
	err := row.Scan(
		&user.ID,
		&user.Login,
		&user.Password,
		&user.IsActive,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
