package main

import (
	"context"
	"net/http"
	"strconv"
	"tiktok/config"
	"tiktok/db/model"
	"tiktok/kitex_gen/favorite"
	"tiktok/service/control/mw"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func favoriteAction(ctx context.Context, c *app.RequestContext) {
	v, ok := c.Get(mw.JWTMiddleware.IdentityKey)
	if !ok {
		hlog.Error(config.NoIDStatusMsg)
		c.JSON(http.StatusBadRequest, &favorite.ActionResponse{
			StatusCode: config.NoIDStatusCode,
			StatusMsg:  &config.NoIDStatusMsg,
		})
		return
	}

	user_id := v.(*model.User).ID
	video_id, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		hlog.Error(err)
		c.JSON(http.StatusBadRequest, &favorite.ActionResponse{
			StatusCode: config.ParameterErrorStatusCode,
			StatusMsg:  &config.ParameterErrorStatusMsg,
		})
		return
	}

	action_type, err := strconv.ParseInt(c.Query("action_type"), 10, 32)
	if err != nil {
		hlog.Error(err)
		c.JSON(http.StatusBadRequest, &favorite.ActionResponse{
			StatusCode: config.ParameterErrorStatusCode,
			StatusMsg:  &config.ParameterErrorStatusMsg,
		})
		return
	}

	hlog.Infof("favoriteAction: user_id: %v, video_id: %v, action_type: %v", user_id, video_id, action_type)

	resp, err := favoriteClient.Action(ctx, &favorite.ActionRequest{
		UserId:     int64(user_id),
		VideoId:    video_id,
		ActionType: int32(action_type),
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

func favoriteList(ctx context.Context, c *app.RequestContext) {
	v, ok := c.Get(mw.JWTMiddleware.IdentityKey)
	if !ok {
		hlog.Error(config.NoIDStatusMsg)
		c.JSON(http.StatusBadRequest, &favorite.ListResponse{
			StatusCode: config.NoIDStatusCode,
			StatusMsg:  &config.NoIDStatusMsg,
		})
		return
	}

	actor_id := v.(*model.User).ID
	user_id, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		hlog.Error(err)
		c.JSON(http.StatusBadRequest, &favorite.ListResponse{
			StatusCode: config.ParameterErrorStatusCode,
			StatusMsg:  &config.ParameterErrorStatusMsg,
		})
		return
	}

	hlog.Infof("favoriteList: user_id: %v, actor_id: %v", user_id, actor_id)

	resp, err := favoriteClient.List(ctx, &favorite.ListRequest{
		ActorId: int64(actor_id),
		UserId:  user_id,
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
