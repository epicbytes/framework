package telemetry

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type SeverityLevel string

func (s SeverityLevel) String() string {
	return string(s)
}

const (
	SeverityTrace SeverityLevel = "TRACE"
	SeverityDebug SeverityLevel = "DEBUG"
	SeverityInfo  SeverityLevel = "INFO"
	SeverityWarn  SeverityLevel = "WARN"
	SeverityError SeverityLevel = "ERROR"
	SeverityFatal SeverityLevel = "FATAL"
	SeverityPanic SeverityLevel = "PANIC"
)

func SendEvent(ctx context.Context, eventName string, severity SeverityLevel, message string, attributes ...attribute.KeyValue) {
	trace.SpanFromContext(ctx).AddEvent(eventName, trace.WithAttributes(
		append(attributes,
			attribute.String("log.severity", severity.String()),
			attribute.String("log.message", message),
		)...,
	))
}

func SendLog(ctx context.Context, severity SeverityLevel, message string, attributes ...attribute.KeyValue) {
	SendEvent(ctx, "log", severity, message, attributes...)
}

func SendLogTrace(ctx context.Context, message string, attributes ...attribute.KeyValue) {
	SendLog(ctx, SeverityTrace, message, attributes...)
}
func SendLogDebug(ctx context.Context, message string, attributes ...attribute.KeyValue) {
	SendLog(ctx, SeverityDebug, message, attributes...)
}
func SendLogInfo(ctx context.Context, message string, attributes ...attribute.KeyValue) {
	SendLog(ctx, SeverityInfo, message, attributes...)
}
func SendLogWarn(ctx context.Context, message string, attributes ...attribute.KeyValue) {
	SendLog(ctx, SeverityWarn, message, attributes...)
}
func SendLogError(ctx context.Context, message string, attributes ...attribute.KeyValue) {
	SendLog(ctx, SeverityError, message, attributes...)
}
func SendLogFatal(ctx context.Context, message string, attributes ...attribute.KeyValue) {
	SendLog(ctx, SeverityFatal, message, attributes...)
}
func SendLogPanic(ctx context.Context, message string, attributes ...attribute.KeyValue) {
	SendLog(ctx, SeverityPanic, message, attributes...)
}

func CreateSpan(ctx context.Context, tracerName string, spanName string) (spanCtx context.Context, span trace.Span) {
	return otel.Tracer(tracerName).Start(ctx, spanName)
}

func ErrorWrapper(ctx context.Context, err error, attributes ...attribute.KeyValue) error {
	SendLogError(ctx, err.Error(), attributes...)
	return err
}

func ErrorWrapperWithSpan(ctx context.Context, err error, span trace.Span, attributes ...attribute.KeyValue) error {
	SendLogError(ctx, err.Error(), attributes...)
	span.RecordError(err, trace.WithStackTrace(true))
	span.SetStatus(codes.Error, err.Error())
	span.SetAttributes(attributes...)
	span.End()
	return err
}
