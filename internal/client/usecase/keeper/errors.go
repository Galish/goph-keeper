package keeper

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrNoConnection    = errors.New("check your connection and try again")
	ErrVersionConflict = errors.New("version conflict")
)

var errorMap = map[codes.Code]error{
	codes.FailedPrecondition: ErrVersionConflict,
	codes.Unavailable:        ErrNoConnection,
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
