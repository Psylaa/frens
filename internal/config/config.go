// internal/config/config.go
package config

import (
	"io/ioutil"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port        string `yaml:"port" validate:"required"`
		JWTSecret   string `yaml:"jwt_secret" validate:"required"`
		JWTDuration int    `yaml:"jwt_duration" validate:"required"`
	} `yaml:"server"`
	Database struct {
		Host     string `yaml:"host" validate:"required"`
		Port     string `yaml:"port" validate:"required"`
		User     string `yaml:"user" validate:"required"`
		DBName   string `yaml:"dbname" validate:"required"`
		Password string `yaml:"password" validate:"required"`
		SSLMode  string `yaml:"sslmode" validate:"required"`
	} `yaml:"database"`
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
