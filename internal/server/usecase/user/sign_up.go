package user

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

func (uc *UserUseCase) SignUp(
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

	err = uc.repo.SetUser(ctx, user)
	if errors.Is(err, repository.ErrConflict) {
		return "", ErrConflict
	}

	if err != nil {
		return "", err
	}

	return uc.jwtManager.Generate(user)
}
