package notes

import (
	"github.com/Galish/goph-keeper/internal/server/repository"
)

type NotesUseCase struct {
	repo repository.Repository
}

func New(repo repository.Repository) *NotesUseCase {
	return &NotesUseCase{
		repo,
	}
}
