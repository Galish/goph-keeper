package notes

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrInvalidEntity   = errors.New("failed entity validation")
	ErrMissingArgument = errors.New("missing required argument")
	ErrNoConnection    = errors.New("check your connection and try again")
	ErrNotFound        = errors.New("entity not found")
	ErrVersionConflict = errors.New("version conflict")
	ErrAuthRequired    = errors.New("authentication required")
)

var errorMap = map[codes.Code]error{
	codes.FailedPrecondition: ErrVersionConflict,
	codes.NotFound:           ErrNotFound,
	codes.Unavailable:        ErrNoConnection,
	codes.Unauthenticated:    ErrAuthRequired,
}

func handleError(err error) error {
	if err == nil {
		return nil
	}

	e, ok := status.FromError(err)
	if !ok {
		return err
	}

	if custom, ok := errorMap[e.Code()]; ok {
		return custom
	}

	return errors.New(e.Message())
}
