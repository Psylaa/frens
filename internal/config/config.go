package config

import (
	"io/ioutil"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
)

type FileType string

const (
	Image FileType = "image"
	Video FileType = "video"
	Audio FileType = "audio"
	Other FileType = "other"
)

type Config struct {
	Server   Server         `yaml:"server" validate:"required"`
	Database Database       `yaml:"database" validate:"required"`
	Storage  StorageDetails `yaml:"storage" validate:"required"`
}

type StorageDetails struct {
	Type  string `yaml:"type" validate:"required"`
	Local struct {
		Path string `yaml:"path" validate:"required"`
	} `yaml:"local"`
	S3 struct {
		Bucket    string `yaml:"bucket" validate:"required"`
		Region    string `yaml:"region" validate:"required"`
		AccessKey string `yaml:"access_key" validate:"required"`
		SecretKey string `yaml:"secret_key" validate:"required"`
	} `yaml:"s3"`
}

type Server struct {
	Port         string `yaml:"port" validate:"required"`
	LogLevel     string `yaml:"log_level" validate:"required"`
	JWTSecret    string `yaml:"jwt_secret" validate:"required"`
	JWTDuration  int    `yaml:"jwt_duration" validate:"required"`
	AllowOrigins string `yaml:"allow_origins" validate:"required"`
}

type Database struct {
	Host         string `yaml:"host" validate:"required"`
	Port         string `yaml:"port" validate:"required"`
	User         string `yaml:"user" validate:"required"`
	DBName       string `yaml:"dbname" validate:"required"`
	Password     string `yaml:"password" validate:"required"`
	SSLMode      string `yaml:"sslmode" validate:"required"`
	LogMode      bool   `yaml:"log_mode"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
}

func (c *Config) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}

func ReadConfig(filename string) (*Config, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		return nil, err
	}

	err = cfg.Validate()
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
