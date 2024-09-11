package session

import (
	"github.com/epicbytes/framework/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis/v3"
	"go.uber.org/fx"
	"sync"
)

var once sync.Once

// NewModule fx module for context.
func NewModule() fx.Option {
	options := fx.Options()
	ctx := func() *session.Store {

		var storage fiber.Storage = nil

		if !helpers.IsDev {
			storage = redis.New(redis.Config{
				Host: "database_redis",
			})
		}

		return session.New(session.Config{
			Storage: storage,
		})
	}
	once.Do(func() {
		options = fx.Options(
			fx.Provide(
				ctx,
			),
		)
	})

	return options
}
