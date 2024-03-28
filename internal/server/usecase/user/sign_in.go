package user

import (
	"context"
	"errors"

	"github.com/Galish/goph-keeper/internal/server/repository"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UserUseCase) SignIn(
	ctx context.Context,
	username, password string,
) (string, error) {
	if username == "" || password == "" {
		return "", ErrMissingCredentials
	}

	user, err := uc.repo.GetUserByLogin(ctx, username)
	if errors.Is(err, repository.ErrUserNotFound) {
		return "", ErrInvalidCredentials
	}

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	return uc.jwtManager.Generate(user)
}