package keeper

import "errors"

var (
	ErrInvalidEntity   = errors.New("failed entity validation")
	ErrInvalidType     = errors.New("invalid entity type")
	ErrMissingArgument = errors.New("missing required argument")
	ErrNotFound        = errors.New("no entity found")
	ErrVersionConflict = errors.New("record version conflict")
	ErrVersionRequired = errors.New("version is required")
)

type VersionError struct {
	version int32
	err     error
}

func NewVersionError(version int32) *VersionError {
	return &VersionError{
		version: version,
		err:     errors.New("record version conflict"),
	}
}

func (e *VersionError) Error() string {
	return e.err.Error()
}
