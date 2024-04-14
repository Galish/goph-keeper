package user

import (
	"context"

	grpc "github.com/Galish/goph-keeper/api/proto"
)

// SignUp implements the creation of a user account.
func (uc *UseCase) SignUp(ctx context.Context, username, password string) error {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	if username == "" || password == "" {
		return ErrInvalidCredentials
	}

	req := &grpc.AuthRequest{
		Username: username,
		Password: password,
	}

	_, err := uc.client.SignUp(ctx, req)

	return handleError(err)
}
