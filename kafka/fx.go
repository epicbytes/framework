package kafka

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"kafka",
		fx.Provide(
			newConfig,
			newKafka,
			fx.Annotate(
				newKafka,
				fx.As(new(Kafka)),
				fx.As(new(KafkaTech)),
			),
		),
		fx.Invoke(func(lc fx.Lifecycle, kf KafkaTech) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						if err := kf.StartKafka(ctx); err != nil {
							kf.Logger().Fatal(fmt.Sprintf("start server error : %v\n", err))
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return kf.StopKafka(ctx)
				},
			})
		}),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("kafka")
		}),
	)
}
