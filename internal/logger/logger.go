package logger

import (
	"os"

	"github.com/bwoff11/frens/internal/config"
	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func Init(level string) {
	logLevel, _ := zerolog.ParseLevel(level)
	zerolog.SetGlobalLevel(logLevel)

	writer := zerolog.ConsoleWriter{Out: os.Stderr}
	Log = zerolog.New(writer).With().Timestamp().Logger()
}

func LogConfig(cfg *config.Config) {
	Log.Info().
		Str("BaseURL", cfg.Server.BaseURL).
		Str("Port", cfg.Server.Port).
		Str("LogLevel", cfg.Server.LogLevel).
		Int("JWTDuration", cfg.Server.JWTDuration).
		Bool("AllowOrigins", cfg.Server.AllowOrigins).
		Msg("Server config loaded")

	Log.Info().
		Str("Host", cfg.Database.Host).
		Str("Port", cfg.Database.Port).
		Str("User", cfg.Database.User).
		Str("DBName", cfg.Database.DBName).
		Str("SSLMode", cfg.Database.SSLMode).
		Bool("LogMode", cfg.Database.LogMode).
		Int("MaxIdleConns", cfg.Database.MaxIdleConns).
		Int("MaxOpenConns", cfg.Database.MaxOpenConns).
		Msg("Database config loaded")

	if cfg.Storage.Type == "local" {
		Log.Info().
			Str("Type", cfg.Storage.Type).
			Str("Path", cfg.Storage.Local.Path).
			Msg("Storage config loaded")
	} else if cfg.Storage.Type == "s3" {
		Log.Info().
			Str("Type", cfg.Storage.Type).
			Str("Bucket", cfg.Storage.S3.Bucket).
			Str("Region", cfg.Storage.S3.Region).
			Str("AccessKey", cfg.Storage.S3.AccessKey).
			Msg("Storage config loaded")
	}

	Log.Info().
		Str("DefaultBio", cfg.Users.DefaultBio).
		Msg("Users config loaded")
}

func DebugLogRequestReceived(pack string, repo string, method string) {
	Log.Debug().
		Str("package", pack).
		Str("repo", repo).
		Str("method", method).
		Msg("request received")
}

func DebugLogRequestUpdate(pack string, repo string, method string, message string) {
	Log.Debug().
		Str("package", pack).
		Str("repo", repo).
		Str("method", method).
		Msgf("request updated: %s", message)
}

func ErrorLogRequestError(pack string, repo string, method string, message string, err error) {
	Log.Error().Err(err).
		Str("package", pack).
		Str("repo", repo).
		Str("method", method).
		Msgf("error handling request: %s", message)
}
