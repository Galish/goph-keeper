package ui

import (
	"io"

	"github.com/fatih/color"
)

func (ui *UI) Print(str string) {
	io.WriteString(ui.w, str)
}

func (ui *UI) Error(err error) {
	color.New(color.FgRed).Fprintf(ui.w, "Error: %s\n", err.Error())
}

func (ui *UI) Break() {
	io.WriteString(ui.w, "\n")
}
