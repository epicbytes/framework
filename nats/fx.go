package nats

import (
	"context"

	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	moduleEntityName = "nats"
)

func NewModule() fx.Option {
	return fx.Module(moduleEntityName,
		fx.Provide(
			fx.Private,
			newConfig,
			newJetStream,
		),
		fx.Invoke(func(lc fx.Lifecycle, cfg *Config, js jetstream.JetStream) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return nil
				},
			})
		}),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named(moduleEntityName)
		}),
	)
}
