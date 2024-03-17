package entity

import "time"

type Note struct {
	ID           string
	Title        string
	Description  string
	Value        string
	RawValue     []byte
	CreatedAt    time.Time
	LastEditedAt time.Time
}
