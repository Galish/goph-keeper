package entity

import (
	"time"

	"github.com/google/uuid"
)

type RawNote struct {
	ID           string
	Title        string
	Description  string
	Value        []byte
	CreatedBy    string
	CreatedAt    time.Time
	LastEditedAt time.Time
}

func NewRawNote() *RawNote {
	return &RawNote{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
	}
}

func (n *RawNote) IsValid() bool {
	return n.Title != "" &&
		n.Value != nil
}
