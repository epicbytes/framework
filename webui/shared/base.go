package shared

import (
	"context"
	b64 "encoding/base64"
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/monoculum/formam/v3"
	"github.com/qor/i18n"
	"net/url"
)

type BaseHtmlTag struct {
	Attrs   templ.Attributes
	Classes templ.Attributes
}

type SpritesBase interface {
	SpritePath() string
}

func GetEntityName(ctx context.Context) string {
	if val, ok := ctx.Value("entityName").(string); ok {
		return val
	}
	return ""
}

func GetRefreshUrl(ctx context.Context) string {
	if val, ok := ctx.Value("entityUrl").(string); ok {
		return val
	}
	return ""
}

func GetEntityUrlPrefix(ctx context.Context) string {
	if val, ok := ctx.Value("entityUrlPrefix").(string); ok {
		return val
	}
	return ""
}

func GetCreateUrl(ctx context.Context) string {
	if val, ok := ctx.Value("entityUrl").(string); ok {
		return val + "/create"
	}
	return ""
}

func GetUpdateUrl(ctx context.Context) string {
	var entityName string
	var entityUrlPrefix string
	var entityId string

	if val, ok := ctx.Value("itemId").(string); ok {
		entityId = val
	}
	if val, ok := ctx.Value("entityName").(string); ok {
		entityName = val
	}
	if val, ok := ctx.Value("entityUrlPrefix").(string); ok {
		entityUrlPrefix = val
	}
	updateUrl, _ := url.JoinPath("/", entityUrlPrefix, entityName, "update", entityId)
	return updateUrl
}

func T(ctx context.Context, scope, key string, args ...interface{}) string {
	if I18n, ok := ctx.Value("i18n").(*i18n.I18n); ok {
		if language, langOk := ctx.Value("currentLanguage").(string); langOk {
			return string(I18n.T(language, scope+"."+key, args...))
		}
	}
	return ""
}

func NotifyData(title, text, status string) map[string]string {
	return map[string]string{
		"type":     "b64",
		"title":    b64.StdEncoding.EncodeToString([]byte(title)),
		"text":     b64.StdEncoding.EncodeToString([]byte(text)),
		"position": "center",
		"status":   status,
	}
}

func ParseFormData[T any](ctx *fiber.Ctx, item T) *fiber.Error {
	if mpd, err := ctx.MultipartForm(); err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	} else {
		decodedValues := formam.NewDecoder(&formam.DecoderOptions{TagName: "form", IgnoreUnknownKeys: true})
		if err = decodedValues.Decode(mpd.Value, item); err != nil {
			return fiber.NewError(fiber.StatusBadGateway, err.Error())
		}
	}
	return nil
}
