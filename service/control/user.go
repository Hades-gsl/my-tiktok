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
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func register(ctx context.Context, c *app.RequestContext) {
	name := c.Query("username")
	password := c.Query("password")

	hlog.Infof("register: name: %v", name)

	resp, err := userClient.Register(ctx, &user.RegisterRequest{
		Username: name,
		Password: password,
	})

	if err, ok := kerrors.FromBizStatusError(err); ok {
		hlog.Error(err.Error())
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status_code": err.BizStatusCode(),
			"status_msg":  err.BizMessage(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func login(ctx context.Context, c *app.RequestContext) {
	name := c.Query("username")
	password := c.Query("password")

	hlog.Infof("login: name: %v", name)

	resp, err := userClient.Login(ctx, &user.LoginRequest{
		Username: name,
		Password: password,
	})

	if err, ok := kerrors.FromBizStatusError(err); ok {
		hlog.Error(err.Error())
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status_code": err.BizStatusCode(),
			"status_msg":  err.BizMessage(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func info(ctx context.Context, c *app.RequestContext) {
	var actor_id int64
	v, ok := c.Get(mw.JWTMiddleware.IdentityKey)
	if !ok {
		hlog.Error("no token or token invalid")
	} else {
		actor_id = int64(v.(float64))
	}
	id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)

	hlog.Infof("info: user_id: %v, actor_id: %v", id, actor_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, &user.InforResponse{
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

	if err, ok := kerrors.FromBizStatusError(err); ok {
		hlog.Error(err.Error())
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status_code": err.BizStatusCode(),
			"status_msg":  err.BizMessage(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
