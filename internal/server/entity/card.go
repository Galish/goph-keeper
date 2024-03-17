package entity

import "time"

type Card struct {
	ID           string
	Title        string
	Description  string
	Number       string
	HolderName   string
	CVC          string
	ExpiryDate   time.Time
	CreatedAt    time.Time
	LastEditedAt time.Time
}
