package entity

import (
	"fmt"
)

type Credentials struct {
	ID          string
	Title       string
	Description string
	Username    string
	Password    string
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
