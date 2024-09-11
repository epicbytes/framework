package postgres

import (
	"fmt"
	"go.uber.org/config"
)

type Config struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SSlMode  string `yaml:"ssl_model"`
}

func newPostgresConfig(provider config.Provider) (*Config, error) {
	var cfg Config
	if err := provider.Get("postgres").Populate(&cfg); err != nil {
		return nil, fmt.Errorf("postgres config: %w", err)
	}
	if cfg.SSlMode == "" {
		cfg.SSlMode = "disable"
	}
	return &cfg, nil
}
