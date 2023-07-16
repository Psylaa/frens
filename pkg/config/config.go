package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type StorageType string

const (
	StorageTypeLocal StorageType = "local"
	StorageTypeS3    StorageType = "s3"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	API      APIConfig      `mapstructure:"handlers"`
	Storage  StorageConfig  `mapstructure:"storage"`
}

type AppConfig struct {
	Users AppUserConfig `mapstructure:"users"`
}

type AppUserConfig struct {
	DefaultBio string `mapstructure:"default_bio"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     string `mapstructure:"port" validate:"required"`
	Name     string `mapstructure:"name" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	LogMode  bool   `mapstructure:"log_mode"`
	SSLMode  string `mapstructure:"ssl_mode"`
	DevMode  bool   `mapstructure:"dev_mode"`
}

type APIConfig struct {
	Port          string `mapstructure:"port" validate:"required"`
	TokenSecret   string `mapstructure:"token_secret" validate:"required"`
	TokenDuration int    `mapstructure:"token_duration" validate:"required"`
}

type StorageConfig struct {
	Type  StorageType        `mapstructure:"type" validate:"required"`
	Local StorageLocalConfig `mapstructure:"local"`
	S3    StorageS3Config    `mapstructure:"s3"`
}

type StorageLocalConfig struct {
	WindowsPath string `mapstructure:"windows_path"`
	LinuxPath   string `mapstructure:"linux_path"`
}

type StorageS3Config struct {
	Bucket    string `mapstructure:"bucket"`
	Region    string `mapstructure:"region"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
}

func (c *Config) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func Load() (*Config, error) {

	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	// Validate
	err := config.Validate()
	if err != nil {
		return nil, err
	}

	return &config, nil
}
