package ui

import (
	"io"
	"os"
)

type UI struct {
	r io.ReadCloser
	w io.WriteCloser
	e io.WriteCloser
}

func New() *UI {
	return &UI{
		r: os.Stdin,
		w: os.Stdout,
		e: os.Stderr,
	}
}

func (ui *UI) Close() error {
	err := ui.r.Close()
	if err != nil {
		return err
	}

	return ui.w.Close()
}
