package notes

import (
	"github.com/Galish/goph-keeper/internal/server/repository"
)

type KeeperUseCase struct {
	repo repository.SecureNotesRepository
}

func New(repo repository.SecureNotesRepository) *KeeperUseCase {
	return &KeeperUseCase{
		repo,
	}
}
