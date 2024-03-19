package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/pkg/auth"
)

func (uc *userUseCase) SignUp(
	ctx context.Context,
	creds *entity.User,
) (string, error) {
	if creds == nil || creds.Login == "" || creds.Password == "" {
		return "", ErrMissingCredentials
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := entity.NewUser()
	user.Login = creds.Login
	user.Password = string(bytes)

	if err := uc.repo.CreateUser(ctx, user); err != nil {
		return "", err
	}

	return auth.GenerateToken(uc.secretKey, user)
}
