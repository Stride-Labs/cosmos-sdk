package log

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Logger defines the interface for logging in the SDK
type Logger interface {
	// TODO define a good interface
}

// Defines commons keys for logging
const ModuleKey = "module"

// ContextKey is used to store the logger in the context
var ContextKey struct{}

func NewZeroLogger(key, value string) Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Kitchen}
	logger := zerolog.New(output).With().Str(key, value).Timestamp().Logger()
	return &logger
}

func NewLogger() Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Kitchen}
	logger := zerolog.New(output).With().Timestamp().Logger()
	return &logger
}
