package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// Interface -.
type Interface interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message interface{}, args ...interface{})
	Fatal(message interface{}, args ...interface{})
}

// Logger -.
type Logger struct {
	logger *zerolog.Logger
}

var (
	_   Interface = (*Logger)(nil)
	Log *Logger
)

// Init -.
func Init(level string) *Logger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	skipFrameCount := 3
	log := zerolog.New(os.Stdout).
		With().
		Timestamp().
		CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).
		Logger()

	Log = new(Logger)
	Log.logger = &log
	return &Logger{logger: &log}
}

// Debug -.
func (l *Logger) Debug(message interface{}, args ...interface{}) {
	l.msg(zerolog.DebugLevel, message, args...)
}

// Info -.
func (l *Logger) Info(message string, args ...interface{}) {
	l.log(zerolog.InfoLevel, message, args...)
}

// Warn -.
func (l *Logger) Warn(message string, args ...interface{}) {
	l.log(zerolog.WarnLevel, message, args...)
}

// Error -.
func (l *Logger) Error(message interface{}, args ...interface{}) {
	if l.logger.GetLevel() == zerolog.DebugLevel {
		l.Debug(message, args...)
	}

	l.msg(zerolog.ErrorLevel, message, args...)
}

// Fatal -.
func (l *Logger) Fatal(message interface{}, args ...interface{}) {
	l.msg(zerolog.FatalLevel, message, args...)

	os.Exit(1)
}

func (l *Logger) log(level zerolog.Level, message string, args ...interface{}) {
	switch level {
	case zerolog.DebugLevel:
		if len(args) == 0 {
			l.logger.Debug().Msg(message)
		} else {
			l.logger.Debug().Msgf(message, args...)
		}
	case zerolog.InfoLevel:
		if len(args) == 0 {
			l.logger.Info().Msg(message)
		} else {
			l.logger.Info().Msgf(message, args...)
		}
	case zerolog.WarnLevel:
		if len(args) == 0 {
			l.logger.Warn().Msg(message)
		} else {
			l.logger.Warn().Msgf(message, args...)
		}
	case zerolog.ErrorLevel:
		if len(args) == 0 {
			l.logger.Error().Msg(message)
		} else {
			l.logger.Error().Msgf(message, args...)
		}
	case zerolog.FatalLevel:
		if len(args) == 0 {
			l.logger.Fatal().Msg(message)
		} else {
			l.logger.Fatal().Msgf(message, args...)
		}
	case zerolog.PanicLevel:
		if len(args) == 0 {
			l.logger.Panic().Msg(message)
		} else {
			l.logger.Panic().Msgf(message, args...)
		}
	}
}

func (l *Logger) msg(level zerolog.Level, message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.log(level, msg.Error(), args...)
	case string:
		l.log(level, msg, args...)
	default:
		l.log(level, fmt.Sprintf("%s message %v has unknown type %v", level, message, msg), args...)
	}
}
