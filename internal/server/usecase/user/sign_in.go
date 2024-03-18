package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/Galish/goph-keeper/pkg/auth"
)

func (uc *userUseCase) SignIn(
	ctx context.Context,
	username, password string,
) (string, error) {
	if username == "" || password == "" {
		return "", ErrMissingCredentials
	}

	user, err := uc.repo.GetByLogin(ctx, username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := auth.GenerateToken(uc.secretKey, user)
	if err != nil {
		return "", err
	}

	return token, nil
}
