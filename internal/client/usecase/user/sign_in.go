package user

import (
	"context"

	grpc "github.com/Galish/goph-keeper/api/proto"
)

func (uc *UserUseCase) SignIn(username, password string) error {
	if username == "" || password == "" {
		return ErrInvalidCredentials
	}

	req := &grpc.AuthRequest{
		Username: username,
		Password: password,
	}

	_, err := uc.client.SignIn(context.Background(), req)

	return handleError(err)
}
