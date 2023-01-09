package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP              `yaml:"http"`
		Log               `yaml:"logger"`
		DB                DB     `yaml:"db"`
		FolderForFiles    string `env-required:"true" yaml:"folder_for_files" env:"FOLDER_FOR_FILES"`
		FolderForArchives string `env-required:"true" yaml:"folder_for_archives" env:"FOLDER_FOR_ARCHIVES"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`
	}

	DB struct {
		Host    string        `env-required:"true" yaml:"host" env:"DB_HOST"`
		Port    string        `env-required:"true" yaml:"port" env:"DB_PORT"`
		User    string        `env-required:"true" yaml:"user" env:"DB_USER"`
		Pass    string        `env-required:"true" yaml:"pass" env:"DB_PASS"`
		Name    string        `env-required:"true" yaml:"name" env:"DB_NAME"`
		Timeout time.Duration `env-default:"500ms" yaml:"timeout" env:"DB_TIMEOUT"`
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
