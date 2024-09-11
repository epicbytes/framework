package otel

import "go.uber.org/zap"

type otel struct {
	config *Config
	logger *zap.Logger
}

func newOtel(config *Config, logger *zap.Logger) *otel {
	return &otel{
		config: config,
		logger: logger,
	}
}
