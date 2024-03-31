package entity

import (
	"fmt"
	"time"
)

const expiryDateLayout = "02/06"

type Card struct {
	ID          string
	Title       string
	Description string
	Number      string
	Holder      string
	CVC         string
	Expiry      time.Time
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
