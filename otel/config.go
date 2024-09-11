package otel

import (
	"fmt"
	"go.uber.org/config"
)

type Config struct {
	Enabled             bool   `yaml:"enabled"`
	CollectorHost       string `yaml:"collector_host"`
	ServiceName         string `yaml:"service_name"`
	InstrumentationName string `yaml:"instrumentation_name"`
	Id                  int64  `yaml:"id"`
	AlwaysOnSampler     bool   `yaml:"always_on_sampler"`
}

func newOtelConfig(provider config.Provider) (*Config, error) {
	var cfg Config
	if err := provider.Get("otel").Populate(&cfg); err != nil {
		return nil, fmt.Errorf("otel config: %w", err)
	}
	return &cfg, nil
}
