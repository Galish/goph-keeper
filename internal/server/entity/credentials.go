package entity

import "time"

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

func (c *Credentials) IsValid() bool {
	return c.Title != "" &&
		c.Username != "" &&
		c.Password != ""
}
