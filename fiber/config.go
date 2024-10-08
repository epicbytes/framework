package fiber

import (
	"fmt"
	"go.uber.org/config"
)

type Config struct {
	Address string `yaml:"address"`
}

func newConfig(provider config.Provider) (*Config, error) {
	var cfg Config
	if err := provider.Get("http_server").Populate(&cfg); err != nil {
		return nil, fmt.Errorf("server config: %w", err)
	}
	return &cfg, nil
}
