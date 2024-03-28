package entity

import (
	"fmt"
	"time"
)

const dateLayout = "02/06"

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
		c.FormatExpiry(c.Expiry),
	)
}

func (c *Card) SetExpiry(input string) error {
	expiry, err := time.Parse(dateLayout, input)
	if err != nil {
		return err
	}

	c.Expiry = expiry

	return nil
}

func (c *Card) FormatExpiry(input time.Time) string {
	return input.Format(dateLayout)
}
