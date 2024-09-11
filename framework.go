package v1

import (
	"github.com/epicbytes/framework/config"
	"github.com/epicbytes/framework/context"
	"github.com/epicbytes/framework/logger"
	"github.com/epicbytes/framework/otel"
	"github.com/epicbytes/framework/vault"
	"go.uber.org/fx"
)

var StandardModules = fx.Options(
	context.NewModule(),
	logger.NewModule(),
	config.NewModule(),
)
var StandardSecurityModules = fx.Options(
	vault.NewModule(),
)
var StandardTraceMetricsModules = fx.Options(
	otel.NewModule(),
)
