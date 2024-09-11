package vault

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"vault",
		fx.Provide(
			newConfig,
			newVault,
		),
		/*fx.Invoke(func(lc fx.Lifecycle, vlt *Vault) {
			lc.Append(fx.Hook{
				OnStart: func(_ context.Context) error {
					//
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return nil
				},
			})
		}),*/
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("vault")
		}),
	)
}
