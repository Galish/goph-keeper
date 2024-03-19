package entity

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID           string
	Title        string
	Description  string
	Value        string
	RawValue     []byte
	CreatedBy    string
	CreatedAt    time.Time
	LastEditedAt time.Time
}

func NewNote() *Note {
	return &Note{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
	}
}

func (n *Note) IsValid() bool {
	return n.Title != "" &&
		(n.Value != "" || n.RawValue != nil)
}
