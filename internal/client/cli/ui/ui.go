package ui

import (
	"io"
	"os"
)

type UserInterface interface {
	Break()
	Confirm(label string) bool
	Edit(string, string, bool) string
	Error(error)
	Input(string, bool) string
	InputPassword(string, bool) string
	Print(string)
	Retry(err error) bool
	Select(string, []*SelectOption)
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
	err := ui.r.Close()
	if err != nil {
		return err
	}

	return ui.w.Close()
}
