package user

import (
	"context"

	grpc "github.com/Galish/goph-keeper/api/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	if err != nil {
		e, ok := status.FromError(err)
		if !ok {
			return err
		}

		switch e.Code() {
		case codes.InvalidArgument:
			return ErrInvalidCredentials

		default:
			return err
		}
	}

	return nil
}
