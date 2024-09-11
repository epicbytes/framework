package bell

import (
	"fmt"
	"go.uber.org/config"
)

type Config struct {
}

func newConfig(provider config.Provider) (*Config, error) {
	var cfg Config
	if err := provider.Get("bell").Populate(&cfg); err != nil {
		return nil, fmt.Errorf("bell config: %w", err)
	}
	return &cfg, nil
}
