package usecase

import (
	"github.com/Galish/goph-keeper/internal/client/entity"
)

type User interface {
	SignUp(string, string) error
	SignIn(string, string) error
}

type Keeper interface {
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

	AddTextNote(creds *entity.TextNote) error
	GetTextNote(id string) (*entity.TextNote, error)
	GetTextNotesList() ([]*entity.TextNote, error)
	UpdateTextNote(creds *entity.TextNote) error
	DeleteTextNote(id string) error

	AddRawNote(creds *entity.RawNote) error
	GetRawNote(id string) (*entity.RawNote, error)
	GetRawNotesList() ([]*entity.RawNote, error)
	UpdateRawNote(creds *entity.RawNote) error
	DeleteRawNote(id string) error
}
