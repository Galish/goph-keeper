package notes

import (
	"errors"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

var (
	ErrInvalidEntity = errors.New("failed entity validation")
	ErrInvalidType   = errors.New("invalid entity type")
)

type NotesUseCase struct {
	repo repository.Repository
}

func New(repo repository.Repository) *NotesUseCase {
	return &NotesUseCase{
		repo,
	}
}
