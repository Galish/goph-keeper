package user

import "errors"

var (
	ErrInvalidCredentials = errors.New("incorrect login/password pair")
	ErrMissingCredentials = errors.New("missing login/password")
	ErrConflict           = errors.New("user already exists")
	ErrNotFound           = errors.New("user not found")
)
