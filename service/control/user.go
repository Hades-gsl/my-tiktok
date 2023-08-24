package main

import (
	"context"
	"net/http"
	"strconv"
	"tiktok/config"
	"tiktok/kitex_gen/user"
	"tiktok/service/control/mw"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func register(ctx context.Context, c *app.RequestContext) {
	hlog.Info("register")
	name := c.Query("username")
	password := c.Query("password")
	resp, err := userClient.Register(ctx, &user.RegisterRequest{
		Username: name,
		Password: password,
	})

	if err != nil {
		hlog.Error(err.Error())
		c.JSON(http.StatusForbidden, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func login(ctx context.Context, c *app.RequestContext) {
	hlog.Info("login")
	name := c.Query("username")
	password := c.Query("password")
	resp, err := userClient.Login(ctx, &user.LoginRequest{
		Username: name,
		Password: password,
	})

	if err != nil {
		hlog.Error(err.Error())
		c.JSON(http.StatusForbidden, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func info(ctx context.Context, c *app.RequestContext) {
	hlog.Info("info")
	actor_id, _ := mw.Auth(c)
	id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusForbidden, &user.InforResponse{
			StatusCode: config.IdInvalidStatusCode,
			StatusMsg:  &config.IdInvalidStatusMsg,
		})
		hlog.Error("id invalid")
		return
	}

	resp, err := userClient.Info(ctx, &user.InfoRequest{
		UserId:  id,
		ActorId: actor_id,
	})

	if err != nil {
		hlog.Error(err.Error())
		c.JSON(http.StatusForbidden, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
