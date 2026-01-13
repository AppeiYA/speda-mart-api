package middlewares

import (
	"context"
	"e-commerce/package/jwt"
)

type contextKey string

const userCtxKey contextKey = "user"

func GetUserFromContext(ctx context.Context) (*jwt.Claims, bool) {
	user, ok := ctx.Value(userCtxKey).(*jwt.Claims)
	return user, ok
}
