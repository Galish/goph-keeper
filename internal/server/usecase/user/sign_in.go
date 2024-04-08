package user

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/Galish/goph-keeper/internal/server/repository"
	"github.com/Galish/goph-keeper/pkg/auth"
)

func (uc *UserUseCase) SignIn(ctx context.Context, username, password string) (string, error) {
	if username == "" || password == "" {
		return "", ErrMissingCredentials
	}

	user, err := uc.repo.GetUserByLogin(ctx, username)
	if errors.Is(err, repository.ErrNotFound) {
		return "", ErrNotFound
	}

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	return uc.jwtManager.Generate(&auth.JWTClaims{
		UserID: user.ID,
	})
}
