package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/pkg/auth"
)

func (uc *userUseCase) SignIn(
	ctx context.Context,
	creds *entity.User,
) (string, error) {
	if creds == nil || creds.Login == "" || creds.Password == "" {
		return "", ErrMissingCredentials
	}

	user, err := uc.repo.GetUserByLogin(ctx, creds.Login)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	return auth.GenerateToken(uc.secretKey, user)
}
