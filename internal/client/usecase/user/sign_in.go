package user

import (
	"context"

	grpc "github.com/Galish/goph-keeper/api/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	if err != nil {
		e, ok := status.FromError(err)
		if !ok {
			return err
		}

		switch e.Code() {
		case codes.InvalidArgument:
			return ErrInvalidCredentials

		case codes.NotFound:
			return ErrNotFound

		default:
			return err
		}
	}

	return nil
}
