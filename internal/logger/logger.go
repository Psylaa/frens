package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger instance
var Logger zerolog.Logger

// LogMessage represents the details for a log message
type LogMessage struct {
	Package  string
	Function string
	Message  string
}

// Init initializes the logger
func Init(level string) {
	logLevel, _ := zerolog.ParseLevel(level)
	zerolog.SetGlobalLevel(logLevel)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

// Debug logs a debug level message
func Debug(logMessage LogMessage) {
	Logger.Debug().
		Str("package", logMessage.Package).
		Str("function", logMessage.Function).
		Msg(logMessage.Message)
}

// Info logs an info level message
func Info(logMessage LogMessage) {
	Logger.Info().
		Str("package", logMessage.Package).
		Str("function", logMessage.Function).
		Msg(logMessage.Message)
}

// Warn logs a warning level message
func Warn(logMessage LogMessage) {
	Logger.Warn().
		Str("package", logMessage.Package).
		Str("function", logMessage.Function).
		Msg(logMessage.Message)
}

// Error logs an error level message
func Error(logMessage LogMessage, err error) {
	Logger.Error().
		Err(err).
		Str("package", logMessage.Package).
		Str("function", logMessage.Function).
		Msg(logMessage.Message)
}

func Fatal(logMessage LogMessage, err error) {
	Logger.Fatal().
		Err(err).
		Str("package", logMessage.Package).
		Str("function", logMessage.Function).
		Msg(logMessage.Message)
}
