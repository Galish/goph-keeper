package cli

import (
	"io"

	"github.com/fatih/color"

	"github.com/Galish/goph-keeper/pkg/logger"
)

func (c *Cli) Print(str string) {
	if _, err := io.WriteString(c.w, str); err != nil {
		logger.WithError(err).Debug("failed writing string")
	}
}

func (c *Cli) Error(err error) {
	color.New(color.FgRed).Fprintf(c.w, "Error: %s\n", err.Error())
}

func (c *Cli) Break() {
	c.Print("\n")
}
