package log

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type Logger struct {
	zerolog.Logger
}

func New() *Logger {
	logger := zerolog.New(os.Stderr)
	return &Logger{logger}
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Error().Msgf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Warn().Msgf(format, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Debug().Msgf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Info().Msgf(format, v...)

}

func NewLog() Logger {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return Logger{logger}
}

//type ContextHook struct{}
//
//func (h *ContextHook) Run(ctx context.Context, e *zerolog.Event) {
//	if traceID := ctx.Value("traceID"); traceID != nil {
//		e.Str("traceID", traceID.(string))
//	}
//}

// Initialize the logger
func Initialize() {
	//zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	//zerolog.LevelFieldName = "level"
	//zerolog.MessageFieldName = "message"
	//zerolog.ErrorFieldName = "error"

}

// Info logs a message at level Info on the global logger.
func Info(ctx context.Context) *zerolog.Event {
	return log.Ctx(ctx).Info()
}

// Debug logs a message at level Debug on the global logger.
func Debug(ctx context.Context) *zerolog.Event {
	return log.Ctx(ctx).Debug()
}

// Warn logs a message at level Warn on the global logger.
func Warn(ctx context.Context) *zerolog.Event {
	return log.Ctx(ctx).Warn()
}

// Error logs a message at level Error on the global logger.
func Error(ctx context.Context) *zerolog.Event {
	return log.Ctx(ctx).Error()
}

// Fatal logs a message at level Fatal on the global logger then the process will exit with status set to 1.
func Fatal(ctx context.Context) *zerolog.Event {
	return log.Ctx(ctx).Fatal()
}

// Panic logs a message at level Panic on the global logger then the process will panic.
func Panic(ctx context.Context) *zerolog.Event {
	return log.Ctx(ctx).Panic()
}

// With returns a child logger with the field added to its context.
func With() zerolog.Context {
	return log.With()
}
