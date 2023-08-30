package mw

import (
	"context"
	"tiktok/db"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

const (
	authResultSuccess string = "success"
	authResultNoToken string = "no_token"
	authResultUnknown string = "unknown"
)

const (
	authResultKey = "authentication_result"
	userIdKey     = "user_id"
)

func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Query("token")
		if token == "" {
			token = string(c.FormValue("token"))
		}

		if token == "" {
			hlog.Info("No token")
			c.Set(authResultKey, authResultNoToken)
			c.Set(userIdKey, 0)
			return
		}

		userToken := db.Q.UserToken
		user, err := userToken.WithContext(ctx).Where(userToken.Token.Eq(token)).First()

		if err != nil {
			hlog.Error(err.Error())
			c.Set(authResultKey, authResultUnknown)
			c.Set(userIdKey, 0)
			return
		}

		c.Set(authResultKey, authResultSuccess)
		c.Set(userIdKey, int64(user.UserID))
	}
}

func Auth(c *app.RequestContext) (id int64, ok bool) {
	result := c.GetString(authResultKey)

	if result == authResultSuccess {
		id = c.GetInt64(userIdKey)
		ok = true
		return
	}

	id = 0
	ok = false
	return
}
