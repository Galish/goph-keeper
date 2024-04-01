package user

import (
	"context"

	grpc "github.com/Galish/goph-keeper/api/proto"
)

func (uc *UserUseCase) SignUp(username, password string) error {
	if username == "" || password == "" {
		return ErrInvalidCredentials
	}

	req := &grpc.AuthRequest{
		Username: username,
		Password: password,
	}

	_, err := uc.client.SignUp(context.Background(), req)

	return handleError(err)
}
