package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

func (uc *userUseCase) SignIn(
	ctx context.Context,
	username, password string,
) (string, error) {
	if username == "" || password == "" {
		return "", ErrMissingCredentials
	}

	user, err := uc.repo.GetUserByLogin(ctx, username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	return uc.jwtManager.Generate(user)
}
