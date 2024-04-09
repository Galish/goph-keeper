package notes

import (
	"github.com/Galish/goph-keeper/internal/server/repository"
)

type UseCase struct {
	repo repository.SecureNotesRepository
}

func New(repo repository.SecureNotesRepository) *UseCase {
	return &UseCase{
		repo,
	}
}
