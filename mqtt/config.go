package mqtt

import (
	"fmt"
	"go.uber.org/config"
)

type Config struct {
	Host          string   `yaml:"host"`
	User          string   `yaml:"user"`
	Password      string   `yaml:"password"`
	ClientId      string   `yaml:"clientId"`
	Subscriptions []string `yaml:"subscriptions"`
}

func newConfig(provider config.Provider) (*Config, error) {
	var cfg Config
	if err := provider.Get("mqtt").Populate(&cfg); err != nil {
		return nil, fmt.Errorf("mqtt config: %w", err)
	}
	return &cfg, nil
}
