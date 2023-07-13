package config

import (
	"github.com/bwoff11/frens/internal/logger"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type FileType string
type StorageType string
type DBType string

const (
	Image    FileType    = "image"
	Video    FileType    = "video"
	Audio    FileType    = "audio"
	Other    FileType    = "other"
	Local    StorageType = "local"
	S3       StorageType = "s3"
	Postgres DBType      = "postgres"
	SQLite   DBType      = "sqlite"
)

type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
	Storage  Storage  `mapstructure:"storage"`
	Users    Users    `mapstructure:"users"`
}

type Server struct {
	BaseURL      string `mapstructure:"base_url"`
	Port         string `mapstructure:"port"`
	LogLevel     string `mapstructure:"log_level"`
	JWTSecret    string `mapstructure:"jwt_secret"`
	JWTDuration  int    `mapstructure:"jwt_duration"`
	AllowOrigins bool   `mapstructure:"allow_origins"`
}

type Database struct {
	Host         string `mapstructure:"host" validate:"required"`
	Port         string `mapstructure:"port" validate:"required"`
	User         string `mapstructure:"user" validate:"required"`
	DBName       string `mapstructure:"dbname" validate:"required"`
	Password     string `mapstructure:"password" validate:"required"`
	SSLMode      string `mapstructure:"sslmode" validate:"required"`
	LogMode      bool   `mapstructure:"log_mode"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

type Storage struct {
	Type  string `mapstructure:"type"`
	Local struct {
		WindowsPath string `mapstructure:"windows_path"`
		LinuxPath   string `mapstructure:"linux_path"`
	} `mapstructure:"local"`
	S3 struct {
		Bucket    string `mapstructure:"bucket"`
		Region    string `mapstructure:"region"`
		AccessKey string `mapstructure:"access_key"`
		SecretKey string `mapstructure:"secret_key"`
	} `mapstructure:"s3"`
}

type Users struct {
	DefaultBio string `mapstructure:"default_bio"`
}

func (c *Config) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func ReadConfig(filename string) (*Config, error) {

	// Set the file name of the configurations file
	viper.SetConfigFile(filename)

	// Manually binding each environment variable
	// AutoEnv wasn't working for some reason, so we may revisit this later
	viper.BindEnv("server.base_url", "BASE_URL")
	viper.BindEnv("server.port", "PORT") // MUST be port for services such as Heroku or Cloud Run
	viper.BindEnv("server.log_level", "LOG_LEVEL")
	viper.BindEnv("server.jwt_secret", "JWT_SECRET")
	viper.BindEnv("server.jwt_duration", "JWT_DURATION")
	viper.BindEnv("server.allow_origins", "ALLOW_ORIGINS")

	viper.BindEnv("database.type", "DB_TYPE")

	viper.BindEnv("database.postgres.host", "DB_HOST")
	viper.BindEnv("database.postgres.port", "DB_PORT")
	viper.BindEnv("database.postgres.dbname", "DB_NAME")
	viper.BindEnv("database.postgres.user", "DB_USER")
	viper.BindEnv("database.postgres.password", "DB_PASSWORD")
	viper.BindEnv("database.postgres.sslmode", "DB_SSLMODE")
	viper.BindEnv("database.postgres.log_mode", "DB_LOG_MODE")
	viper.BindEnv("database.postgres.max_idle_conns", "DB_MAX_IDLE_CONNS")
	viper.BindEnv("database.postgres.max_open_conns", "DB_MAX_OPEN_CONNS")

	viper.BindEnv("database.sqlite.path", "SQLITE_PATH")
	viper.BindEnv("database.sqlite.log_mode", "SQLITE_LOG_MODE")
	viper.BindEnv("database.sqlite.max_idle_conns", "SQLITE_MAX_IDLE_CONNS")
	viper.BindEnv("database.sqlite.max_open_conns", "SQLITE_MAX_OPEN_CONNS")

	viper.BindEnv("storage.type", "STORAGE_TYPE")
	viper.BindEnv("storage.local.path", "LOCAL_PATH")
	viper.BindEnv("storage.s3.bucket", "S3_BUCKET")
	viper.BindEnv("storage.s3.region", "S3_REGION")
	viper.BindEnv("storage.s3.access_key", "S3_ACCESS_KEY")
	viper.BindEnv("storage.s3.secret_key", "S3_SECRET_KEY")

	viper.BindEnv("users.default_bio", "DEFAULT_BIO")

	// Try to read the config file and output any encountered error
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Once the config file has been read, unmarshal it to the Config struct
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	// Validate if the Config struct was populated correctly
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	// If everything went well, return the Config struct
	return &cfg, nil
}

func (cfg *Config) Print() {
	logger.Log.Info().
		Str("BaseURL", cfg.Server.BaseURL).
		Str("Port", cfg.Server.Port).
		Str("LogLevel", cfg.Server.LogLevel).
		Int("JWTDuration", cfg.Server.JWTDuration).
		Bool("AllowOrigins", cfg.Server.AllowOrigins).
		Msg("Server config loaded")

	logger.Log.Info().
		Str("Host", cfg.Database.Host).
		Str("Port", cfg.Database.Port).
		Str("User", cfg.Database.User).
		Str("DBName", cfg.Database.DBName).
		Str("SSLMode", cfg.Database.SSLMode).
		Bool("LogMode", cfg.Database.LogMode).
		Int("MaxIdleConns", cfg.Database.MaxIdleConns).
		Int("MaxOpenConns", cfg.Database.MaxOpenConns).
		Msg("Postgres database config loaded")

	if cfg.Storage.Type == "local" {
		logger.Log.Info().
			Str("Type", cfg.Storage.Type).
			Str("WindowsPath", cfg.Storage.Local.WindowsPath).
			Str("LinuxPath", cfg.Storage.Local.LinuxPath).
			Msg("Local storage config loaded")
	} else if cfg.Storage.Type == "s3" {
		logger.Log.Info().
			Str("Type", cfg.Storage.Type).
			Str("Bucket", cfg.Storage.S3.Bucket).
			Str("Region", cfg.Storage.S3.Region).
			Str("AccessKey", cfg.Storage.S3.AccessKey).
			Msg("S3 storage config loaded")
	}

	logger.Log.Info().
		Str("DefaultBio", cfg.Users.DefaultBio).
		Msg("Users config loaded")
}
