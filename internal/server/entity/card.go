package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

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

func NewCard() *Card {
	return &Card{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
	}
}

func (c *Card) IsValid() bool {
	fmt.Println("+Card+", c, c.Title != "",
		c.Number != "",
		c.Holder != "",
		c.CVC != "",
		!c.Expiry.IsZero(),
	)

	return c.Title != "" &&
		c.Number != "" &&
		c.Holder != "" &&
		c.CVC != "" &&
		!c.Expiry.IsZero()
}
