package entity

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Credentials struct {
	ID           string
	Title        string
	Description  string
	Username     string
	Password     string
	CreatedBy    string
	CreatedAt    time.Time
	LastEditedAt time.Time
}

func NewCredentials() *Credentials {
	return &Credentials{
		ID:        uuid.NewString(),
		CreatedAt: time.Now(),
	}
}

func (c *Credentials) IsValid() bool {
	return c.Title != "" &&
		c.Username != "" &&
		c.Password != ""
}

func (c *Credentials) String() string {
	return fmt.Sprintf(
		"Title: %s\nDescription: %s\nUsername: %s\nPassword: %s\n",
		c.Title,
		c.Description,
		c.Username,
		c.Password,
	)
}
