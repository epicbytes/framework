package postgres

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewModule() fx.Option {
	return fx.Module(
		"postgres",
		fx.Provide(
			newPostgresConfig,
			newPostgres,
		),
		fx.Invoke(func(lc fx.Lifecycle, pg *Postgres, logger *zap.Logger) {
			lc.Append(fx.Hook{
				OnStart: func(_ context.Context) error {
					go func() {
						if err := pg.startMigrations(); err != nil {
							logger.Error("start postgres server error", zap.Error(err))
						}
					}()
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return pg.Conn.Close()
				},
			})
		}),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named("postgres")
		}),
	)
}
