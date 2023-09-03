package mw

import (
	"context"
	"log"
	"tiktok/db/model"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"gorm.io/gorm"
)

//when token is necessary, we use this middleware.
var identityKey = "identity_key"
var err error

var JWTMiddleware *jwt.HertzJWTMiddleware

func init() {
	JWTMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:       "tiktok",
		Key:         []byte("tiktok secret key"),
		Timeout:     time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &model.User{
				Model: gorm.Model{
					ID: claims[identityKey].(uint),
				},
			}
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"status_code": code,
				"status_msg":  message,
			})
		},
		TokenLookup:   "query: token, form: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
}
