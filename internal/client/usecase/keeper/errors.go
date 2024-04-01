package keeper

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrAlreadyExists      = errors.New("user already exists")
	// ErrInvalidCredentials = errors.New("incorrect login/password pair")

	ErrNoConnection = errors.New("check your connection and try again")
	// Error: rpc error: code = InvalidArgument desc = failed entity validation

	// ErrNoPassword         = errors.New("password not specified")
	// ErrNoUsername         = errors.New("username not specified")
	// ErrNotFound           = errors.New("user not found")
)

var errorMap = map[codes.Code]error{
	// codes.InvalidArgument: ErrInvalidCredentials,
	// codes.NotFound:        ErrNotFound,
	// codes.AlreadyExists:   ErrAlreadyExists,
	codes.Unavailable: ErrNoConnection,
}

func handleError(err error) error {
	if err == nil {
		return nil
	}

	e, ok := status.FromError(err)
	if !ok {
		return err
	}

	custom, ok := errorMap[e.Code()]
	if !ok {
		return err
	}

	return custom
}
