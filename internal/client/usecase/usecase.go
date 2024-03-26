package usecase

import (
	"github.com/Galish/goph-keeper/internal/client/entity"
)

type User interface {
	SignUp(string, string) error
	SignIn(string, string) error
}

type Keeper interface {
	// AddNote(ctx context.Context, note *entity.Note) error
	// AddCard(ctx context.Context, card *entity.Card) error
	// AddCredentials(ctx context.Context, creds *entity.Credentials) error

	// GetNote(ctx context.Context, user, id string) (*entity.Note, error)
	// GetCard(ctx context.Context, user, id string) (*entity.Card, error)
	GetCredentials(id string) (*entity.Credentials, error)

	// GetTextNotes(ctx context.Context, user string) ([]*entity.Note, error)
	// GetRawNotes(ctx context.Context, user string) ([]*entity.Note, error)
	// GetCards(ctx context.Context, user string) ([]*entity.Card, error)
	GetAllCredentials() ([]*entity.Credentials, error)

	DeleteCredentials(id string) error
}
