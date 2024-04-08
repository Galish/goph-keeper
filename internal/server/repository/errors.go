package repository

import "errors"

var (
	ErrConflict        = errors.New("entity id already exists")
	ErrNotFound        = errors.New("nothing was found")
	ErrVersionConflict = errors.New("entity version conflict")
)
