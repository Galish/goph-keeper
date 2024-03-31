package entity

import (
	"fmt"
)

type TextNote struct {
	ID          string
	Title       string
	Description string
	Value       string
}

func (n *TextNote) String() string {
	return fmt.Sprintf(
		"Title: %s\nDescription: %s\nNote: %s\n",
		n.Title,
		n.Description,
		n.Value,
	)
}
