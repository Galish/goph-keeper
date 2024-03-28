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

	AddCard(card *entity.Card) error
	GetCard(id string) (*entity.Card, error)
	GetCardsList() ([]*entity.Card, error)
	UpdateCard(creds *entity.Card) error
	DeleteCard(id string) error

	AddCredentials(creds *entity.Credentials) error
	GetCredentials(id string) (*entity.Credentials, error)
	GetCredentialsList() ([]*entity.Credentials, error)
	UpdateCredentials(creds *entity.Credentials) error
	DeleteCredentials(id string) error

	// GetNote(ctx context.Context, user, id string) (*entity.Note, error)
	// GetCard(ctx context.Context, user, id string) (*entity.Card, error)

	// GetTextNotes(ctx context.Context, user string) ([]*entity.Note, error)
	// GetRawNotes(ctx context.Context, user string) ([]*entity.Note, error)

}
