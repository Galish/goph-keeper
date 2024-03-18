package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/Galish/goph-keeper/pkg/auth"
)

func (uc *userUseCase) SignUp(
	ctx context.Context,
	username, password string,
) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user, err := uc.repo.CreateUser(ctx, username, string(bytes))
	if err != nil {
		return "", err
	}

	token, err := auth.GenerateToken(uc.secretKey, user)
	if err != nil {
		return "", err
	}

	return token, nil
}
