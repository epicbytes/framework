package grpc

import (
	"fmt"

	"go.uber.org/config"
)

type Config struct {
	Endpoint        string `yaml:"endpoint"`
	AllowReflection bool   `yaml:"allow_reflection"`
}

func newConfig(provider config.Provider) (*Config, error) {
	var cfg Config
	if err := provider.Get(moduleEntityName).Populate(&cfg); err != nil {
		return nil, fmt.Errorf("%s config: %w", moduleEntityName, err)
	}
	return &cfg, nil
}
