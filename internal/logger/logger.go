package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func Init(level string) {
	logLevel, _ := zerolog.ParseLevel(level)
	zerolog.SetGlobalLevel(logLevel)

	writer := zerolog.ConsoleWriter{Out: os.Stderr}
	Log = zerolog.New(writer).With().Timestamp().Logger()
}
