package fiber

import (
	"context"

	"github.com/goccy/go-json"
	"github.com/gofiber/template/html/v2"
	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type MiddlewareHandlerMap map[string][]fiber.Handler

type Params struct {
	fx.In
	Engine       *html.Engine          `optional:"true"` // template generation engine
	ErrorHandler *fiber.ErrorHandler   `optional:"true"` // custom error handler
	Middleware   *MiddlewareHandlerMap `optional:"true"` // custom middleware
}

func newFiber(params Params) *fiber.App {
	var cfg fiber.Config
	var app *fiber.App

	var errorHandler func(*fiber.Ctx, error) error

	if params.ErrorHandler != nil {
		errorHandler = *params.ErrorHandler

	} else {
		errorHandler = func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
	}

	cfg = fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: errorHandler,
	}

	if params.Engine != nil {
		cfg.Views = params.Engine
	}

	app = fiber.New(cfg)

	if params.Middleware != nil {
		for path, handlers := range *params.Middleware {
			for _, handler := range handlers {
				if handler == nil {
					continue
				}
				if path == "" {
					app.Use(handler)
				} else {
					app.Use(path, handler)
				}
			}
		}
	}

	return app
}

func useFiber(
	lifecycle fx.Lifecycle,
	app *fiber.App,
	logger *zap.Logger,
	cfg *Config,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := app.Listen(cfg.Address); err != nil {
					logger.Fatal(err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := app.Shutdown(); err != nil {
				logger.Fatal(err.Error())
			}
			return nil
		},
	})
}
