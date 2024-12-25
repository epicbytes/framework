package kit

import (
	"errors"
	"strings"

	"buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	"connectrpc.com/connect"
	"github.com/gofiber/fiber/v2"
)

func InjectValidationError(ctx *fiber.Ctx, error error, addons ...map[string]string) {
	viols := make(map[string]string)
	for _, addon := range addons {
		for k, v := range addon {
			viols[k] = v
		}
	}
	var connectError *connect.Error
	if errors.As(error, &connectError) {
		for _, errr := range connectError.Details() {
			viol, _ := errr.Value()
			if v := viol.(*validate.Violations); v != nil {
				for _, vv := range v.Violations {
					viols[strings.Replace(strings.TrimPrefix(*vv.FieldPath, "item."), "].", "]", -1)] = *vv.Message
				}
			}
		}
	}
	ctx.Context().SetUserValue("errors", viols)
}
