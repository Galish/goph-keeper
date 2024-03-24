package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"

	"github.com/Galish/goph-keeper/internal/server/entity"
)

func (uc *userUseCase) SignUp(
	ctx context.Context,
	username, password string,
) (string, error) {
	if username == "" || password == "" {
		return "", ErrMissingCredentials
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := entity.NewUser()
	user.Login = username
	user.Password = string(bytes)

	if err := uc.repo.SetUser(ctx, user); err != nil {
		return "", err
	}

	return uc.jwtManager.Generate(user)
}
