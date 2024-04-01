package repository

import "errors"

var (
	ErrNotFound = errors.New("nothing was found")
	ErrConflict = errors.New("record id already exists")
)
