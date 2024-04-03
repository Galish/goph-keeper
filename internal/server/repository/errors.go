package repository

import "errors"

var (
	ErrConflict        = errors.New("record id already exists")
	ErrNotFound        = errors.New("nothing was found")
	ErrVersionConflict = errors.New("record version conflict")
)
