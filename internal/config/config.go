package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// FileType is a type that represents the different types of file we might deal with
type FileType string

const (
	Image FileType = "image"
	Video FileType = "video"
	Audio FileType = "audio"
	Other FileType = "other"
)

// Config struct stores all configuration of our application
// It will be populated either from a config file or environment variables
type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
	Storage  Storage  `mapstructure:"storage"`
	Users    Users    `mapstructure:"users"`
}

// Server struct represents the server details of our application
type Server struct {
	BaseURL      string `mapstructure:"base_url"`
	Port         string `mapstructure:"port"`
	LogLevel     string `mapstructure:"log_level"`
	JWTSecret    string `mapstructure:"jwt_secret"`
	JWTDuration  int    `mapstructure:"jwt_duration"`
	AllowOrigins bool   `mapstructure:"allow_origins"`
}

// Database struct represents the database details of our application
type Database struct {
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	User         string `mapstructure:"user"`
	DBName       string `mapstructure:"dbname"`
	Password     string `mapstructure:"password"`
	SSLMode      string `mapstructure:"sslmode"`
	LogMode      bool   `mapstructure:"log_mode"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
}

// StorageDetails struct represents the storage details of our application
type Storage struct {
	Type  string `mapstructure:"type"`
	Local struct {
		Path string `mapstructure:"path"`
	} `mapstructure:"local"`
	S3 struct {
		Bucket    string `mapstructure:"bucket"`
		Region    string `mapstructure:"region"`
		AccessKey string `mapstructure:"access_key"`
		SecretKey string `mapstructure:"secret_key"`
	} `mapstructure:"s3"`
}

// Users struct represents the users' default details of our application
type Users struct {
	DefaultBio string `mapstructure:"default_bio"`
}

// Validate method validates if the Config struct is properly populated
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

	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.port", "DB_PORT")
	viper.BindEnv("database.name", "DB_NAME")
	viper.BindEnv("database.user", "DB_USER")
	viper.BindEnv("database.password", "DB_PASSWORD")
	viper.BindEnv("database.sslmode", "SSL_MODE")
	viper.BindEnv("database.log_mode", "LOG_MODE")
	viper.BindEnv("database.max_idle_conns", "MAX_IDLE_CONNS")
	viper.BindEnv("database.max_open_conns", "MAX_OPEN_CONNS")

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
