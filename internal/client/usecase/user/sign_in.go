package user

import (
	"context"

	grpc "github.com/Galish/goph-keeper/api/proto"
)

func (uc *UserUseCase) SignIn(ctx context.Context, username, password string) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	if username == "" || password == "" {
		return ErrInvalidCredentials
	}

	req := &grpc.AuthRequest{
		Username: username,
		Password: password,
	}

	_, err := uc.client.SignIn(ctx, req)

	return handleError(err)
}
