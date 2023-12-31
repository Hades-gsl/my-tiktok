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
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func feedList(ctx context.Context, c *app.RequestContext) {
	v, ok := c.Get(mw.JWTMiddleware.IdentityKey)
	var id int64
	if ok {
		id = int64(v.(float64))
	}

	latestTime, err := strconv.ParseInt(c.Query("LatestTime"), 10, 64)
	if err != nil {
		hlog.Error(err)
		latestTime = time.Now().Unix()
	}

	hlog.Infof("feedList: id: %v, latestTime: %v", id, latestTime)

	res, err := feedClient.List(ctx, &feed.ListRequest{
		LatestTime: &latestTime,
		UserId:     &id,
	})

	if err, ok := kerrors.FromBizStatusError(err); ok {
		hlog.Error(err.Error())
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status_code": err.BizStatusCode(),
			"status_msg":  err.BizMessage(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
