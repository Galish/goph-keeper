// Package logger implements a simple logger.
package logger

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

var logger = &log.Logger{
	Formatter: new(log.JSONFormatter),
	Out:       io.Discard,
	Level:     log.DebugLevel,
}

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

// Init initializes the logger.
func Init() {
	logger.SetOutput(os.Stderr)
}

// SetLevel sets the logger level.
func SetLevel(level string) {
	logLevel, err := log.ParseLevel(level)
	if err != nil {
		logger.Error("invalid log level", err)
		return
	}

	logger.SetLevel(logLevel)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	logger.Log(log.DebugLevel, args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	logger.Log(log.InfoLevel, args...)
}

// WithError creates an entry from the standard logger and adds an error to it.
func WithError(err error) *log.Entry {
	return log.NewEntry(logger).WithError(err)
}

// WithFields creates an entry from the standard logger and adds multiple fields to it.
func WithFields(fields Fields) *log.Entry {
	f := log.Fields{}

	for k, v := range fields {
		f[k] = v
	}

	return log.NewEntry(logger).WithFields(f)
}
