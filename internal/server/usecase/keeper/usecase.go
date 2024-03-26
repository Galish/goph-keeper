package keeper

import (
	"errors"

	"github.com/Galish/goph-keeper/internal/server/repository"
)

var (
	ErrInvalidEntity = errors.New("failed entity validation")
	ErrInvalidType   = errors.New("invalid entity type")
)

type KeeperUseCase struct {
	repo repository.KeeperRepository
}

func New(repo repository.KeeperRepository) *KeeperUseCase {
	return &KeeperUseCase{
		repo,
	}
}
