package entity

import "time"

type Credentials struct {
	ID           string
	Title        string
	Description  string
	Username     string
	Password     string
	CreatedAt    time.Time
	LastEditedAt time.Time
}
