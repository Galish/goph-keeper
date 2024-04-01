package keeper

import (
	"github.com/Galish/goph-keeper/internal/server/repository"
)

type KeeperUseCase struct {
	repo repository.KeeperRepository
}

func New(repo repository.KeeperRepository) *KeeperUseCase {
	return &KeeperUseCase{
		repo,
	}
}
