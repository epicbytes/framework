package otel

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"otel",
		fx.Provide(
			newOtelConfig,
			//NewOpa,
		),
		fx.Invoke(func(lc fx.Lifecycle, config *Config) {
			lc.Append(fx.Hook{
				OnStart: func(_ context.Context) error {
					return nil
				},
			})
		}),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("otel")
		}),
	)
}
