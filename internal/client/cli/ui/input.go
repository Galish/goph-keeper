package ui

import (
	"bufio"
	"fmt"
	"strings"
)

func (ui *UI) Input(label string, isRequired bool) string {
	var input string

	r := bufio.NewReader(ui.r)

	for {
		ui.Print(fmt.Sprintf("%s: ", label))

		input, _ = r.ReadString('\n')
		input = strings.TrimSpace(input)

		if !isRequired || input != "" {
			break
		}
	}

	return input
}
