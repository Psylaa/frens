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

func DebugLogRequestRecieved(pack string, repo string, method string) {
	Log.Debug().
		Str("package", pack).
		Str("repo", repo).
		Str("method", method).
		Msg("request received")
}

func ErrorLogRequestError(pack string, repo string, method string, err error) {
	Log.Error().Err(err).
		Str("package", pack).
		Str("repo", repo).
		Str("method", method).
		Msg("error handling request")
}
