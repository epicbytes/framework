package telemetry

import (
	"fmt"
	"go.uber.org/config"
)

type Config struct {
	OtelCollectorUrl string `yaml:"otel_collector_url"`
	OtelServiceName  string `yaml:"otel_service_name"`
}

func newConfig(provider config.Provider) (*Config, error) {
	var cfg Config
	if err := provider.Get("tracing").Populate(&cfg); err != nil {
		return nil, fmt.Errorf("%s config: %w", "tracing", err)
	}
	return &cfg, nil
}
