package telemetry

import (
	"context"
	"fmt"
	otel2 "github.com/agoda-com/opentelemetry-logs-go"
	"github.com/agoda-com/opentelemetry-logs-go/exporters/otlp/otlplogs"
	"github.com/agoda-com/opentelemetry-logs-go/sdk/logs"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initConnection(config *Config) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(config.OtelCollectorUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC connection to collector: %w", err)
	}

	return conn, nil
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTraceProvider(config *Config, ctx context.Context, conn *grpc.ClientConn) (*trace.TracerProvider, error) {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(config.OtelServiceName),
		),
	)

	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create trace exporter: %w", err)
	}

	tp := trace.NewTracerProvider(
		trace.WithResource(r),
		trace.WithSpanProcessor(trace.NewBatchSpanProcessor(
			traceExporter,
		)),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

func newMeterProvider(config *Config, ctx context.Context, conn *grpc.ClientConn) (*metric.MeterProvider, error) {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(config.OtelServiceName),
		),
	)
	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
	if err != nil {
		return nil, fmt.Errorf("failed to create metric exporter: %w", err)
	}
	meterProvider := metric.NewMeterProvider(
		metric.WithResource(r),
		metric.WithReader(
			metric.NewPeriodicReader(metricExporter),
		),
	)
	otel.SetMeterProvider(meterProvider)

	return meterProvider, nil
}

func newLoggerProvider(config *Config, ctx context.Context) (*logs.LoggerProvider, error) {

	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(config.OtelServiceName),
		),
	)

	logExporter, err := otlplogs.NewExporter(ctx)
	if err != nil {
		return nil, err
	}

	loggerProvider := logs.NewLoggerProvider(
		logs.WithBatcher(logExporter),
		logs.WithResource(r),
	)
	otel2.SetLoggerProvider(loggerProvider)
	return loggerProvider, nil
}
