package entity

import (
	"fmt"
)

type RawNote struct {
	ID          string
	Title       string
	Description string
	Value       []byte
}

func (n *RawNote) String() string {
	return fmt.Sprintf(
		"Title: %s\nDescription: %s\nNote: %s\n",
		n.Title,
		n.Description,
		n.Value,
	)
}
