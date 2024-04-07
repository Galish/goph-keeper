package repository

import (
	"context"

	"github.com/Galish/goph-keeper/internal/entity"
)

type UserRepository interface {
	AddUser(context.Context, *entity.User) error
	GetUserByLogin(context.Context, string) (*entity.User, error)
}
