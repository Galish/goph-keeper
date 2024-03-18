package repository

import (
	"context"

	"github.com/Galish/goph-keeper/internal/server/entity"
)

type UserRepository interface {
	CreateUser(context.Context, string, string) (*entity.User, error)
	GetUserByLogin(context.Context, string) (*entity.User, error)
}
