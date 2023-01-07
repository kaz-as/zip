package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP   `yaml:"http"`
		Log    `yaml:"logger"`
		DB     DB     `yaml:"db"`
		Folder string `env-required:"true" yaml:"folder" env:"FOLDER"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	DB struct {
		Host string `env-required:"true" yaml:"host" env:"DB_HOST"`
		Port string `env-required:"true" yaml:"port" env:"DB_PORT"`
		User string `env-required:"true" yaml:"user" env:"DB_USER"`
		Pass string `env-required:"true" yaml:"pass" env:"DB_PASS"`
		Name string `env-required:"true" yaml:"name" env:"DB_NAME"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
