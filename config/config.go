package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
)

func New(opts ...Option) *Config {
	cfg := &Config{}

	for _, o := range opts {
		o.apply(cfg)
	}

	if err := env.Parse(cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return cfg
}
