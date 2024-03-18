package user

import (
	"errors"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

var ErrMissingCredentials = errors.New("missing login/password")
var ErrInvalidCredentials = errors.New("incorrect login/password pair")

type userUseCase struct {
	repo      repository.UserRepository
	secretKey string
}

func New(repo repository.UserRepository, secretKey string) *userUseCase {
	return &userUseCase{
		repo:      repo,
		secretKey: secretKey,
	}
}
