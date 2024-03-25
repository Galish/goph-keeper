package user

import (
	"errors"

	"github.com/Galish/goph-keeper/internal/server/repository"
	"github.com/Galish/goph-keeper/pkg/auth"
)

var (
	ErrMissingCredentials = errors.New("missing login/password")
	ErrInvalidCredentials = errors.New("incorrect login/password pair")
)

type UserUseCase struct {
	repo       repository.UserRepository
	jwtManager *auth.JWTManager
}

func New(repo repository.UserRepository, jwtManager *auth.JWTManager) *UserUseCase {
	return &UserUseCase{
		repo:       repo,
		jwtManager: jwtManager,
	}
}
