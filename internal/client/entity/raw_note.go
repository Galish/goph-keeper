package entity

import (
	"fmt"
	"strconv"
	"strings"
)

type RawNote struct {
	ID          string
	Title       string
	Description string
	Value       []byte
}

func (n *RawNote) String() string {
	return fmt.Sprintf(
		"Title: %s\nDescription: %s\nNote: %s\n",
		n.Title,
		n.Description,
		n.GetValue(),
	)
}

func (n *RawNote) GetValue() string {
	return strings.Trim(fmt.Sprintf("%v", n.Value), "[]")
}

func (n *RawNote) SetValue(input string) error {
	str := strings.Split(input, " ")
	rawValue := make([]byte, len(str))

	for i, s := range str {
		parsed, err := strconv.ParseInt(s, 10, 8)
		if err != nil {
			return err
		}

		rawValue[i] = byte(parsed)
	}

	n.Value = rawValue

	return nil
}
