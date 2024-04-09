package ui

import (
	"io"

	"github.com/Galish/goph-keeper/pkg/logger"
	"github.com/fatih/color"
)

func (ui *UI) Print(str string) {
	if _, err := io.WriteString(ui.w, str); err != nil {
		logger.WithError(err).Debug("failed writing string")
	}
}

func (ui *UI) Error(err error) {
	color.New(color.FgRed).Fprintf(ui.w, "Error: %s\n", err.Error())
}

func (ui *UI) Break() {
	ui.Print("\n")
}
