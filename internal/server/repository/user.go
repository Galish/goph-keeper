package repository

import (
	"context"
	"errors"

	"github.com/Galish/goph-keeper/internal/server/entity"
)

var (
	ErrUserConflict = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository interface {
	SetUser(context.Context, *entity.User) error
	GetUserByLogin(context.Context, string) (*entity.User, error)
}
