package auth

import (
	"connectrpc.com/connect"
	"context"
)

const TFAForwardHeader = "X-Forwarded-User"
const TFAForwardHeaderToken = "X-Forwarded-Token"
const UserDataHeader = "X-User-Data"
const UserContextName = "userData"
const TokenContextName = "token"

func NewAuthPropagatorInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				if user, ok := ctx.Value(UserContextName).(*BaseUserModel); ok {
					req.Header().Set(UserDataHeader, user.Encode())
				}
			}
			if !req.Spec().IsClient && len(req.Header().Get(UserDataHeader)) > 0 {
				var decodedUser = new(BaseUserModel)
				err := decodedUser.Decode(req.Header().Get(UserDataHeader))
				if err != nil {
					return nil, err
				}
				authCtx := context.WithValue(ctx, UserContextName, decodedUser)
				return next(authCtx, req)
			}
			return next(ctx, req)
		}
	}
}
