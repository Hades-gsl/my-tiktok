package main

import (
	"context"
	"net/http"
	"strconv"
	"tiktok/config"
	"tiktok/kitex_gen/publish"
	"tiktok/service/control/mw"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func publishAction(ctx context.Context, c *app.RequestContext) {
	id, ok := mw.Auth(c)
	if !ok {
		hlog.Error(config.TokenInvalidStatusMsg)
		c.JSON(http.StatusForbidden, &publish.ActionResponse{
			StatusCode: config.TokenInvalidStatusCode,
			StatusMsg:  &config.TokenInvalidStatusMsg,
		})
		return
	}

	title := c.FormValue("title")
	file, err := c.FormFile("data")
	if err != nil {
		hlog.Error(err)
		c.JSON(http.StatusForbidden, &publish.ActionResponse{
			StatusCode: config.GetFileErrorStatusCode,
			StatusMsg:  &config.GetFileErrorStatusMsg,
		})
		return
	}

	fp, err := file.Open()
	if err != nil {
		hlog.Error(err)
		c.JSON(http.StatusForbidden, &publish.ActionResponse{
			StatusCode: config.FileOpenErrorStatusCode,
			StatusMsg:  &config.FileOpenErrorStatusMsg,
		})
		return
	}
	defer func() {
		if err := fp.Close(); err != nil {
			hlog.Error(err)
		}
	}()

	data := make([]byte, file.Size)
	size, err := fp.Read(data)
	if err != nil || size != int(file.Size) {
		if err != nil {
			hlog.Error(err)
		} else {
			hlog.Error(config.FileReadErrorStatusMsg)
		}
		c.JSON(http.StatusForbidden, &publish.ActionResponse{
			StatusCode: config.FileReadErrorStatusCode,
			StatusMsg:  &config.FileReadErrorStatusMsg,
		})
		return
	}

	hlog.Infof("publishAction: user_id: %v, title: %v", id, string(title))

	resp, err := publishClinet.Action(ctx, &publish.ActionRequest{
		UserId: id,
		Data:   data,
		Title:  string(title),
	})

	if err != nil {
		hlog.Error(err)
		c.JSON(http.StatusForbidden, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func publishList(ctx context.Context, c *app.RequestContext) {
	actor_id, ok := mw.Auth(c)
	if !ok {
		hlog.Error(config.TokenInvalidStatusMsg)
		c.JSON(http.StatusForbidden, &publish.ListResponse{
			StatusCode: config.TokenInvalidStatusCode,
			StatusMsg:  &config.TokenInvalidStatusMsg,
		})
		return
	}

	user_id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		hlog.Error(err)
		c.JSON(http.StatusForbidden, &publish.ListResponse{
			StatusCode: config.IdInvalidStatusCode,
			StatusMsg:  &config.IdInvalidStatusMsg,
		})
		return
	}

	hlog.Infof("publishList: user_id: %v, actor_id: %v", user_id, actor_id)

	resp, err := publishClinet.List(ctx, &publish.ListRequest{
		UserId:  user_id,
		ActorId: actor_id,
	})

	if err != nil {
		hlog.Error(err)
		c.JSON(http.StatusForbidden, resp)
		return
	}

	c.JSON(http.StatusOK, resp)
}
