package user

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/Galish/goph-keeper/internal/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
	"github.com/Galish/goph-keeper/pkg/auth"
)

func (uc *UseCase) SignUp(ctx context.Context, username, password string) (string, error) {
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

	err = uc.repo.AddUser(ctx, user)
	if errors.Is(err, repository.ErrConflict) {
		return "", ErrConflict
	}

	if err != nil {
		return "", err
	}

	return uc.jwtManager.Generate(&auth.JWTClaims{
		UserID: user.ID,
	})
}
