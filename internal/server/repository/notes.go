package repository

import (
	"context"
	"time"
)

type SecureNotesRepository interface {
	AddSecureNote(ctx context.Context, note *SecureNote) error
	UpdateSecureNote(ctx context.Context, note *SecureNote) error
	DeleteSecureNote(ctx context.Context, user, id string, noteType SecureNoteType) error
	GetSecureNote(ctx context.Context, user, id string, noteType SecureNoteType) (*SecureNote, error)
	GetSecureNotes(ctx context.Context, user string, noteType SecureNoteType) ([]*SecureNote, error)
}

const (
	TypeCredentials SecureNoteType = iota + 1
	TypeTextNote
	TypeRawNote
	TypeCard
)

type SecureNoteType int

type SecureNote struct {
	ID           string
	Type         SecureNoteType
	Title        string
	Description  string
	Username     string
	Password     string
	TextNote     string
	RawNote      []byte
	CardNumber   string
	CardHolder   string
	CardCVC      string
	CardExpiry   time.Time
	CreatedBy    string
	CreatedAt    time.Time
	LastEditedAt time.Time
	Version      int32
}
