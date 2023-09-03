package mw

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/form3tech-oss/jwt-go"
)

//when token is optional, we use this middleware.
func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Query("token")
		if token == "" {
			token = string(c.FormValue("token"))
		}

		if token == "" {
			hlog.Info("No token")
			return
		}

		t, err := JWTMiddleware.ParseTokenString(token)
		if err != nil {
			hlog.Error(err)
			return
		}

		claims := t.Claims.(jwt.MapClaims)
		c.Set(identityKey, claims[identityKey])
	}
}
