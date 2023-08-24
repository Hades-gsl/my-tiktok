package main

import (
	"context"
	"net/http"
	"strconv"
	"tiktok/kitex_gen/feed"
	"tiktok/service/control/mw"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func feedList(ctx context.Context, c *app.RequestContext) {
	hlog.Info("feedList")
	id, _ := mw.Auth(c)
	latestTime, err := strconv.ParseInt(c.Query("LatestTime"), 10, 64)
	if err != nil {
		hlog.Error(err.Error())
		latestTime = time.Now().Unix()
	}

	res, err := feedClient.List(ctx, &feed.ListRequest{
		LatestTime: &latestTime,
		UserId:     &id,
	})

	if err != nil {
		hlog.Error(err.Error())
	}

	c.JSON(http.StatusOK, res)
}
