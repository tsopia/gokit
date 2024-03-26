package log

import (
	"context"
	"github.com/rs/zerolog"
	"os"
)

var Logger *zerolog.Logger

func InitLog() {
	l := zerolog.New(os.Stderr).With().Timestamp().Logger()
	//l.Level(zerolog.DebugLevel)
	Logger = &l
}

func init() {
	InitLog()
}

// Initialize the logger
func Initialize() {
	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	//zerolog.LevelFieldName = "level"
	//zerolog.MessageFieldName = "message"
	//zerolog.ErrorFieldName = "error"

}

// Info logs a message at level Info on the global logger.
func Info(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Info()
}

// Debug logs a message at level Debug on the global logger.
func Debug(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Debug()
}

// Warn logs a message at level Warn on the global logger.
func Warn(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Warn()
}

// Error logs a message at level Error on the global logger.
func Error(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Error()
}

// Fatal logs a message at level Fatal on the global logger then the process will exit with status set to 1.
func Fatal(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Fatal()
}

// Panic logs a message at level Panic on the global logger then the process will panic.
func Panic(ctx context.Context) *zerolog.Event {
	return Ctx(ctx).Panic()
}

// With returns a child logger with the field added to its context.
func With() zerolog.Context {
	return Logger.With()
}

// Ctx returns a contextual logger. If the context is empty, the global logger is returned.
func Ctx(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}
