package keeper

import "errors"

var (
	ErrInvalidEntity = errors.New("failed entity validation")
	ErrInvalidType   = errors.New("invalid entity type")
	ErrNotFound      = errors.New("no entity found")
)
