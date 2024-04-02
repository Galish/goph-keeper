package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string
	Login     string
	Password  string
	IsActive  bool
	CreatedAt time.Time
}

func NewUser() *User {
	return &User{
		ID:        uuid.NewString(),
		IsActive:  true,
		CreatedAt: time.Now(),
	}
}
