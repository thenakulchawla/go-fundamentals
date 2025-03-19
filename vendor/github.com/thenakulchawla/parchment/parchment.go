// Package parchment for contextual logger
package parchment

import (
	"context"
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

// Define a unique type for the logger context key
type contextKey struct{}

// Key used to store the logger in the context
var Key = contextKey{}

// LoggerField represents a structured logging field (key-value pair)
type LoggerField struct {
	Key   string
	Value any
}

// New initializes a new logger and stores it in the context
func New(ctx context.Context) context.Context {
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()

	return context.WithValue(ctx, Key, logger)
}

// AddToLogger adds multiple structured fields (key-value pairs) to the logger
func AddToLogger(ctx context.Context, fields []LoggerField) context.Context {
	logger := FromContext(ctx)
	newLoggerContext := logger.With()

	for _, field := range fields {
		switch v := field.Value.(type) {
		case string:
			newLoggerContext = newLoggerContext.Str(field.Key, v)
		case int:
			newLoggerContext = newLoggerContext.Int(field.Key, v)
		case float64:
			newLoggerContext = newLoggerContext.Float64(field.Key, v)
		case bool:
			newLoggerContext = newLoggerContext.Bool(field.Key, v)
		default:
			// Convert other types to string
			newLoggerContext = newLoggerContext.Str(field.Key, fmt.Sprintf("%v", v))
		}
	}

	newLogger := newLoggerContext.Logger()
	return context.WithValue(ctx, Key, newLogger)
}

// FromContext retrieves the logger from the context, or returns a no-op logger if not found
func FromContext(ctx context.Context) zerolog.Logger {
	logger, ok := ctx.Value(Key).(zerolog.Logger)
	if !ok {
		return zerolog.Nop() // No-op logger to avoid nil dereference
	}
	return logger
}
