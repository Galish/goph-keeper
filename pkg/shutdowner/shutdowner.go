// Package shutdowner implements a utility to perform application graceful shutdown.
package shutdowner

import (
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/Galish/goph-keeper/pkg/logger"
)

// Shutdowner implements graceful shutdown.
type Shutdowner struct {
	closers []func() error
	done    chan struct{}
}

// New returns a new Shutdowner instance.
func New(cs ...io.Closer) *Shutdowner {
	sd := Shutdowner{
		done: make(chan struct{}),
	}

	sd.Add(cs...)

	sigs := make(chan os.Signal, 1)

	signal.Notify(
		sigs,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)

	go func() {
		<-sigs
		sd.Shutdown()
	}()

	return &sd
}

// Add saves items to close on shutdown.
func (sd *Shutdowner) Add(cs ...io.Closer) {
	for _, c := range cs {
		sd.closers = append(sd.closers, c.Close)
	}
}

// Shutdown closes all items.
func (sd *Shutdowner) Shutdown() {
	for _, c := range sd.closers {
		if err := c(); err != nil {
			logger.WithError(err).Debug("failed to close")
		}
	}

	close(sd.done)
}

// Wait for the shutdown to complete.
func (sd *Shutdowner) Wait() {
	<-sd.done
}
