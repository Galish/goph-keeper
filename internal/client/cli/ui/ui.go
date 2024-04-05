package ui

import (
	"io"
	"os"

	"github.com/Galish/goph-keeper/pkg/logger"
)

type UserInterface interface {
	Break()
	Confirm(string) bool
	Edit(string, string, bool) string
	Error(error)
	Input(string, bool) string
	InputPassword(string, bool) string
	Print(string)
	Retry(error) bool
	Select(string, []*SelectOption)
	ReadFile(string, bool) []byte
	WriteFile(string, []byte, bool)
}

type UI struct {
	r io.ReadCloser
	w io.WriteCloser
}

func New() *UI {
	return &UI{
		r: os.Stdin,
		w: os.Stdout,
	}
}

func (ui *UI) Close() error {
	logger.Info("shutting down the CLI application")

	if err := ui.r.Close(); err != nil {
		return err
	}

	return ui.w.Close()
}
