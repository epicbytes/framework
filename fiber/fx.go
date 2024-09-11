package fiber

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"fiber",
		fx.Provide(newConfig, newFiber),
		fx.Invoke(useFiber),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("fiber")
		}),
	)
}
