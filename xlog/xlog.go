// xlog/xlog.go
package xlog

import (
	"github.com/rs/zerolog"
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
