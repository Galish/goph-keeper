package entity

import (
	"fmt"
	"reflect"
	"time"
)

type Credentials struct {
	ID           string
	Title        string
	Description  string
	Username     string
	Password     string
	CreatedAt    time.Time
	LastEditedAt time.Time
}

func (c *Credentials) Print() string {
	var str string

	c.Iterate(func(key, value string) {
		switch key {
		case "Title", "Description", "Username", "Password":
			str += fmt.Sprintf("%s: %s\n", key, value)
		}
	})

	return str
}

func (c *Credentials) Iterate(fn func(string, string)) {
	values := reflect.ValueOf(*c)
	types := values.Type()

	for i := 0; i < values.NumField(); i++ {
		fn(types.Field(i).Name, values.Field(i).String())
	}
}
