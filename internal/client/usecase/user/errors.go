package user

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrAlreadyExists      = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("incorrect login/password pair")
	ErrNoConnection       = errors.New("check your connection and try again")
	ErrNotFound           = errors.New("user not found")
)

var errorMap = map[codes.Code]error{
	codes.InvalidArgument: ErrInvalidCredentials,
	codes.NotFound:        ErrNotFound,
	codes.AlreadyExists:   ErrAlreadyExists,
	codes.Unavailable:     ErrNoConnection,
}

func handleError(err error) error {
	if err == nil {
		return nil
	}

	e, ok := status.FromError(err)
	if !ok {
		return err
	}

	if e.Code() == codes.Internal {
		return errors.New(e.Message())
	}

	custom, ok := errorMap[e.Code()]
	if !ok {
		return err
	}

	return custom
}
