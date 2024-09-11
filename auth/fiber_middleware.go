package auth

import (
	"bytes"
	"github.com/gofiber/contrib/opafiber/v2"
	"github.com/gofiber/fiber/v2"
)

func Middleware(map[string][]PolicyKey) func(app *fiber.Ctx) error {
	module := `
package example.authz

default allow := false

allow {
	true
}
`

	cfg := opafiber.Config{
		RegoQuery:             "data.example.authz.allow",
		RegoPolicy:            bytes.NewBufferString(module),
		IncludeQueryString:    true,
		DeniedStatusCode:      fiber.StatusForbidden,
		DeniedResponseMessage: "status forbidden",
		IncludeHeaders:        []string{"Authorization"},
		InputCreationMethod: func(ctx *fiber.Ctx) (map[string]interface{}, error) {
			return map[string]interface{}{
				"method": ctx.Method(),
				"path":   ctx.Path(),
			}, nil
		},
	}
	return opafiber.New(cfg)
}
