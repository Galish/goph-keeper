package usecase

import (
	"context"

	"github.com/Galish/goph-keeper/internal/entity"
)

type User interface {
	SignUp(context.Context, string, string) (token string, err error)
	SignIn(context.Context, string, string) (token string, err error)
	Verify(accessToken string) (*entity.User, error)
}

type Keeper interface {
	AddCredentials(ctx context.Context, creds *entity.Credentials) error
	GetCredentials(ctx context.Context, user, id string) (*entity.Credentials, error)
	GetAllCredentials(ctx context.Context, user string) ([]*entity.Credentials, error)
	UpdateCredentials(ctx context.Context, creds *entity.Credentials) error
	DeleteCredentials(ctx context.Context, user, id string) error

	AddCard(ctx context.Context, card *entity.Card) error
	GetCard(ctx context.Context, user, id string) (*entity.Card, error)
	GetCards(ctx context.Context, user string) ([]*entity.Card, error)
	UpdateCard(ctx context.Context, card *entity.Card) error
	DeleteCard(ctx context.Context, user, id string) error

	AddTextNote(ctx context.Context, note *entity.TextNote) error
	GetTextNote(ctx context.Context, user, id string) (*entity.TextNote, error)
	GetTextNotes(ctx context.Context, user string) ([]*entity.TextNote, error)
	UpdateTextNote(ctx context.Context, note *entity.TextNote) error
	DeleteTextNote(ctx context.Context, user, id string) error

	AddRawNote(ctx context.Context, note *entity.RawNote) error
	GetRawNote(ctx context.Context, user, id string) (*entity.RawNote, error)
	GetRawNotes(ctx context.Context, user string) ([]*entity.RawNote, error)
	UpdateRawNote(ctx context.Context, note *entity.RawNote) error
	DeleteRawNote(ctx context.Context, user, id string) error
}
