package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type TextNote struct {
	ID           string
	Title        string
	Description  string
	Value        string
	CreatedBy    string
	CreatedAt    time.Time
	LastEditedAt time.Time
}

func NewTextNote() *TextNote {
	return &TextNote{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
	}
}

func (n *TextNote) IsValid() bool {
	return n.Title != "" &&
		n.Value != ""
}

func (n *TextNote) String() string {
	return fmt.Sprintf(
		"Title: %s\nDescription: %s\nNote: %s\n",
		n.Title,
		n.Description,
		n.Value,
	)
}
