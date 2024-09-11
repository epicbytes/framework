package bell

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"sync"
)

var once sync.Once

func NewModule() fx.Option {
	options := fx.Options()
	once.Do(func() {
		options = fx.Module(
			"bell",
			fx.Provide(
				newConfig,
				newBell,
			),
			fx.Invoke(func(lc fx.Lifecycle) {
				lc.Append(fx.Hook{
					OnStart: func(_ context.Context) error {
						return nil
					},
					OnStop: func(ctx context.Context) error {
						return nil
					},
				})
			}),
			fx.Decorate(func(log *zap.Logger) *zap.Logger {
				return log.Named("bell")
			}),
		)
	})
	return options
}
