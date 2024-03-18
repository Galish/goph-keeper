package entity

import "time"

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

func (n *Note) IsValid() bool {
	return n.Title != "" &&
		(n.Value != "" || n.RawValue != nil)
}
