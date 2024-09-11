package auth

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

func PushUserDataToContext(ctx *fiber.Ctx) {
	userDataRaw := ctx.Get(TFAForwardHeader)
	if userDataRaw == "" {
		return
	}
	var user = new(BaseUserModel)
	user.Decode(userDataRaw)
	if user == nil {
		return
	}

	ctx.Context().SetUserValue(UserContextName, user)
	ctx.SetUserContext(context.WithValue(ctx.UserContext(), UserContextName, user))

	token := ctx.Get(TFAForwardHeaderToken)
	ctx.Context().SetUserValue(TokenContextName, token)
	ctx.SetUserContext(context.WithValue(ctx.UserContext(), TokenContextName, token))
}

func GetUser(ctx context.Context) *BaseUserModel {
	user, ok := ctx.Value(UserContextName).(*BaseUserModel)
	if !ok {
		return nil
	}
	return user
}

func GetUserFullName(ctx context.Context) string {
	if user := GetUser(ctx); user != nil {
		return user.Name
	}
	return ""
}

func GetUserEmail(ctx context.Context) string {
	if user := GetUser(ctx); user != nil {
		return user.Email
	}
	return ""
}

func GetUserResourceAccess(ctx context.Context, resourceAccess interface{}) error {
	if user := GetUser(ctx); user != nil {
		b, err := json.Marshal(user.ResourceAccessRaw)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, resourceAccess)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}
