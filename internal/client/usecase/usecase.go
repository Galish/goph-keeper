package usecase

import (
	"github.com/Galish/goph-keeper/internal/entity"
)

type User interface {
	SignUp(string, string) error
	SignIn(string, string) error
}

type Keeper interface {
	AddCard(*entity.Card) error
	GetCard(string) (*entity.Card, error)
	GetCardsList() ([]*entity.Card, error)
	UpdateCard(*entity.Card, bool) error
	DeleteCard(string) error

	AddCredentials(*entity.Credentials) error
	GetCredentials(string) (*entity.Credentials, error)
	GetCredentialsList() ([]*entity.Credentials, error)
	UpdateCredentials(*entity.Credentials, bool) error
	DeleteCredentials(string) error

	AddTextNote(*entity.TextNote) error
	GetTextNote(string) (*entity.TextNote, error)
	GetTextNotesList() ([]*entity.TextNote, error)
	UpdateTextNote(*entity.TextNote, bool) error
	DeleteTextNote(string) error

	AddRawNote(*entity.RawNote) error
	GetRawNote(string) (*entity.RawNote, error)
	GetRawNotesList() ([]*entity.RawNote, error)
	UpdateRawNote(*entity.RawNote, bool) error
	DeleteRawNote(string) error
}
