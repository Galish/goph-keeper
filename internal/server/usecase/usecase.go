package usecase

import (
	"context"

	"github.com/Galish/goph-keeper/internal/server/entity"
)

type User interface {
	SignUp(context.Context, string, string) (token string, err error)
	SignIn(context.Context, string, string) (token string, err error)
	Verify(string) (*entity.User, error)
}

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

	DeleteTextNote(ctx context.Context, user, id string) error
	DeleteRawNote(ctx context.Context, user, id string) error
	DeleteCard(ctx context.Context, user, id string) error
	DeleteCredentials(ctx context.Context, user, id string) error

	UpdateCredentials(ctx context.Context, creds *entity.Credentials) error
}
