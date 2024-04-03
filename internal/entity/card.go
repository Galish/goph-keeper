package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const expiryDateLayout = "02/06"

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
	Version      int32
}

func NewCard() *Card {
	return &Card{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
	}
}

func (c *Card) IsValid() bool {
	return c.Title != "" &&
		c.Number != "" &&
		c.Holder != "" &&
		c.CVC != "" &&
		!c.Expiry.IsZero()
}

func (c *Card) String() string {
	return fmt.Sprintf(
		"Title: %s\nDescription: %s\nCard number: %s\nCard holder: %s\nCVC: %s\nExpiration date: %s\n",
		c.Title,
		c.Description,
		c.Number,
		c.Holder,
		c.CVC,
		c.GetExpiry(),
	)
}

func (c *Card) GetExpiry() string {
	return c.Expiry.Format(expiryDateLayout)
}

func (c *Card) SetExpiry(input string) error {
	expiry, err := time.Parse(expiryDateLayout, input)
	if err != nil {
		return err
	}

	c.Expiry = expiry

	return nil
}
