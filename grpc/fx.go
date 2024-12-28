package grpc

import (
	"context"
	"net"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	moduleEntityName = "grpc"
)

func NewModule() fx.Option {
	return fx.Module(moduleEntityName,
		fx.Provide(
			fx.Private,
			newConfig,
		),
		fx.Invoke(func(lc fx.Lifecycle, cfg *Config, logger *zap.Logger, s *grpc.Server) {
			lc.Append(fx.Hook{
				OnStart: func(_ context.Context) error {
					lis, err := net.Listen("tcp", cfg.Endpoint)
					if err != nil {
						return err
					}

					// DEVELOPMENT only !
					if cfg.AllowReflection {
						reflection.Register(s)
					}

					go func() {
						if err := s.Serve(lis); err != nil {
							logger.Error("error", zap.Error(err))
							return
						}
					}()

					return nil
				},
				OnStop: func(ctx context.Context) error {
					s.GracefulStop()
					return nil
				},
			})
		}),
		fx.Decorate(func(log *zap.Logger) *zap.Logger {
			return log.Named(moduleEntityName)
		}),
	)
}
