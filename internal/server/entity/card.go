package entity

import "time"

type Card struct {
	ID           string
	Title        string
	Description  string
	Number       string
	Holder       string
	CVC          string
	Expiry       time.Time
	CreatedBy    string
	CreatedAt    time.Time
	LastEditedAt time.Time
}

func (c *Card) IsValid() bool {
	return c.Title != "" &&
		c.Number != "" &&
		c.Holder != "" &&
		c.CVC != "" &&
		!c.Expiry.IsZero()
}
