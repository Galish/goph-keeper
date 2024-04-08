package usecase

import (
	"context"

	"github.com/Galish/goph-keeper/internal/entity"
)

type User interface {
	SignUp(context.Context, string, string) error
	SignIn(context.Context, string, string) error
}

type SecureNotes interface {
	AddCard(context.Context, *entity.Card) error
	GetCard(context.Context, string) (*entity.Card, error)
	GetCardsList(context.Context) ([]*entity.Card, error)
	UpdateCard(context.Context, *entity.Card, bool) error
	DeleteCard(context.Context, string) error

	AddCredentials(context.Context, *entity.Credentials) error
	GetCredentials(context.Context, string) (*entity.Credentials, error)
	GetCredentialsList(context.Context) ([]*entity.Credentials, error)
	UpdateCredentials(context.Context, *entity.Credentials, bool) error
	DeleteCredentials(context.Context, string) error

	AddTextNote(context.Context, *entity.TextNote) error
	GetTextNote(context.Context, string) (*entity.TextNote, error)
	GetTextNotesList(context.Context) ([]*entity.TextNote, error)
	UpdateTextNote(context.Context, *entity.TextNote, bool) error
	DeleteTextNote(context.Context, string) error

	AddRawNote(context.Context, *entity.RawNote) error
	GetRawNote(context.Context, string) (*entity.RawNote, error)
	GetRawNotesList(context.Context) ([]*entity.RawNote, error)
	UpdateRawNote(context.Context, *entity.RawNote, bool) error
	DeleteRawNote(context.Context, string) error
}
