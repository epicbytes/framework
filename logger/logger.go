package logger

import (
	"github.com/mattn/go-colorable"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newLogger() (*otelzap.Logger, *zap.Logger, error) {
	lg := zap.NewDevelopmentEncoderConfig()
	lg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger := otelzap.New(zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(lg),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	)), otelzap.WithCaller(false), otelzap.WithMinLevel(zap.DebugLevel))
	zap.ReplaceGlobals(logger.Logger)
	otelzap.ReplaceGlobals(logger)

	otel.SetErrorHandler(otel.ErrorHandlerFunc(func(err error) {
		logger.Error("error", zap.Error(err))
	}))

	return logger, logger.Logger, nil
}
