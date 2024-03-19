package keeper

import (
	"context"
	"errors"

	"github.com/Galish/goph-keeper/internal/server/entity"
	"github.com/Galish/goph-keeper/internal/server/repository"
)

var (
	ErrInvalidEntity = errors.New("failed entity validation")
	ErrInvalidType   = errors.New("invalid entity type")
)

type Keeper interface {
	AddNote(ctx context.Context, note *entity.Note) error
	AddCard(ctx context.Context, card *entity.Card) error
	AddCredentials(ctx context.Context, creds *entity.Credentials) error

	GetNote(ctx context.Context, user, id string) (*entity.Note, error)
	GetCard(ctx context.Context, user, id string) (*entity.Card, error)
	GetCredentials(ctx context.Context, user, id string) (*entity.Credentials, error)

	GetTextNotes(ctx context.Context, user string) ([]*entity.Note, error)
	GetRawNotes(ctx context.Context, user string) ([]*entity.Note, error)
	GetCards(ctx context.Context, user string) ([]*entity.Card, error)
	GetAllCredentials(ctx context.Context, user string) ([]*entity.Credentials, error)
}

type KeeperUseCase struct {
	repo repository.KeeperRepository
}

func New(repo repository.KeeperRepository) *KeeperUseCase {
	return &KeeperUseCase{
		repo,
	}
}
