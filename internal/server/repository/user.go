package repository

import (
	"context"

	"github.com/Galish/goph-keeper/internal/server/entity"
)

type UserRepository interface {
	Create(context.Context, string, string) (*entity.User, error)
	GetByLogin(context.Context, string) (*entity.User, error)
}
