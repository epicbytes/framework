package telemetry

import (
	"context"
	sdk "github.com/agoda-com/opentelemetry-logs-go/sdk/logs"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type Result struct {
	fx.Out

	Mp *metric.MeterProvider
	Tp *trace.TracerProvider
	Lp *sdk.LoggerProvider
	Pp propagation.TextMapPropagator
}

func NewModule() fx.Option {
	return fx.Module(
		"tracing",
		fx.Provide(
			fx.Private,
			newConfig,
		),
		fx.Provide(
			initConnection,
			newTracer,
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
	)
}

func newTracer(config *Config, ctx context.Context, conn *grpc.ClientConn) (Result, error) {
	tp, err := newTraceProvider(config, ctx, conn)
	if err != nil {
		return Result{}, err
	}
	mp, err := newMeterProvider(config, ctx, conn)
	if err != nil {
		return Result{}, err
	}
	lp, err := newLoggerProvider(config, ctx)
	if err != nil {
		return Result{}, err
	}
	
	return Result{
		Mp: mp,
		Tp: tp,
		Lp: lp,
		Pp: newPropagator(),
	}, nil
}
