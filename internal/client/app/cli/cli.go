package cli

import (
	"io"
	"os"
	"syscall"

	"github.com/Galish/goph-keeper/pkg/logger"
)

type UserInterface interface {
	Break()
	Close() error
	Confirm(string) bool
	Edit(string, string, bool) string
	Error(error)
	Exit()
	Input(string, bool) string
	InputPassword(string, bool) string
	Print(string)
	ReadFile(string, bool) []byte
	Retry(error) bool
	Select(string, []*SelectOption)
	WriteFile(string, []byte, bool)
}

type Cli struct {
	r io.ReadCloser
	w io.WriteCloser
}

func New() *Cli {
	return &Cli{
		r: os.Stdin,
		w: os.Stdout,
	}
}

func (c *Cli) Exit() {
	if err := syscall.Kill(syscall.Getpid(), syscall.SIGINT); err != nil {
		logger.WithError(err).Debug("failed exit")
	}
}

func (c *Cli) Close() error {
	logger.Info("shutting down the CLI application")

	if err := c.r.Close(); err != nil {
		return err
	}

	return c.w.Close()
}
