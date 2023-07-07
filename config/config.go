package config

import (
	"github.com/caarlos0/env/v8"
	"github.com/rs/zerolog/log"
)

func New(opts ...Option) *Config {
	cfg := &Config{}

	for _, o := range opts {
		o.apply(cfg)
	}

	if err := env.Parse(cfg); err != nil {
		log.Error().Msgf("%+v\n", err)
	}

	return cfg
}
