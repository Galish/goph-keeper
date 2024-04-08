package notes

import "errors"

var (
	ErrInvalidEntity   = errors.New("failed entity validation")
	ErrInvalidType     = errors.New("invalid entity type")
	ErrMissingArgument = errors.New("missing required argument")
	ErrNotFound        = errors.New("no entity found")
	ErrVersionConflict = errors.New("note version conflict")
	ErrVersionRequired = errors.New("version is required")
)
