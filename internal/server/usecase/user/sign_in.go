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
	user, err := uc.repo.GetUserByLogin(ctx, username)
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
