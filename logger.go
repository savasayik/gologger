// Module: github.com/savasayik/gologger

package gologger

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// --- LogLevel Definition ---
type LogLevel string

const (
	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
	FatalLevel LogLevel = "fatal"
)

// --- Logger Definition ---
type Logger[T any] struct {
	base zerolog.Logger
	tag  string
}

var logInstance *Logger[any]

// InitLogger initializes the global logger.
func InitLogger(level LogLevel, tag string) {
	zerolog.TimeFieldFormat = time.RFC3339
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("tag", tag).Logger()

	switch level {
	case DebugLevel:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case InfoLevel:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case WarnLevel:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case ErrorLevel:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case FatalLevel:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	logInstance = &Logger[any]{base: logger, tag: tag}
}

// GetLogger returns the global logger instance.
func GetLogger() *Logger[any] {
	if logInstance == nil {
		panic("logger not initialized")
	}
	return logInstance
}

// WithContext adds contextual fields from context.Context.
func (l *Logger[T]) WithContext(ctx context.Context) *zerolog.Event {
	e := l.base.Info()
	if traceID, ok := ctx.Value("trace_id").(string); ok {
		e = e.Str("trace_id", traceID)
	}
	if userID, ok := ctx.Value("user_id").(string); ok {
		e = e.Str("user_id", userID)
	}
	if reqID, ok := ctx.Value("request_id").(string); ok {
		e = e.Str("request_id", reqID)
	}
	return e
}

// WithErrorStack logs error with stack trace.
func (l *Logger[T]) WithErrorStack(err error, msg string) {
	stacked := fmt.Sprintf("%+v", err)
	l.base.Error().Str("error", err.Error()).Str("stack", stacked).Msg(msg)
}

// StructuredLog logs typed payload with level and event name.
func (l *Logger[T]) StructuredLog(level string, event string, data T) {
	e := l.base.WithLevel(parseLevel(level))
	e.Str("event", event).Interface("payload", data).Msg("structured log")
}

// StructuredDebug logs at debug level.
func (l *Logger[T]) StructuredDebug(event string, data T) {
	l.StructuredLog("debug", event, data)
}

// StructuredError logs at error level.
func (l *Logger[T]) StructuredError(event string, data T) {
	l.StructuredLog("error", event, data)
}

// --- Helpers ---

func parseLevel(level string) zerolog.Level {
	switch level {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	default:
		return zerolog.InfoLevel
	}
}
