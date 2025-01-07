package http_server

import (
	v1 "github.com/epicbytes/framework"
	"github.com/epicbytes/framework/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func New(ready chan struct{}) *fx.App {
	app := fx.New(
		fx.Options(
			v1.StandardModules,
			NewModule(),
			fx.Provide(DummyErrorHandler),
		),
		logger.Decorate(),
	)
	close(ready)
	return app
}

var DummyErrorHandler = func(*fiber.Ctx, error) error {
	return nil
}
