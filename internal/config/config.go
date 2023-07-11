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

// ReadConfig function reads configuration from a file or environment variables and validates it
func ReadConfig(filename string) (*Config, error) {

	// Set the file name of the configurations file
	viper.SetConfigFile(filename)

	// Set default values for the database configs
	// They will be used if no other value is provided through file or environment variable
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
	viper.SetDefault("database.user", "sampleuser")
	viper.SetDefault("database.password", "pass")

	// viper.AutomaticEnv() configures viper to read from environment variables
	// An environment variable with name 'X' will be matched with a key 'x' in Viper
	viper.AutomaticEnv()

	// Try to read the config file and output any encountered error
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Once the config file has been read, unmarshal it to the Config struct
	// UnmarshalExact requires all fields in the Config struct to be present in the config file
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
