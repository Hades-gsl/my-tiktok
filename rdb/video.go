package rdb

import (
	"context"
	"strconv"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func GetLikesAndCommentsCount(ctx context.Context, id int64) (likesCount, commentsCount int64, err error) {
	var fc, cc string
	fc, err = RedisDB.Get(ctx, strconv.FormatInt(id, 10)+"_likes").Result()
	if err != nil {
		hlog.Error(err)
		return
	}

	cc, err = RedisDB.Get(ctx, strconv.FormatInt(id, 10)+"_comments").Result()
	if err != nil {
		hlog.Error(err)
		return
	}

	likesCount, err = strconv.ParseInt(fc, 10, 64)
	if err != nil {
		hlog.Error(err)
		return
	}

	commentsCount, err = strconv.ParseInt(cc, 10, 64)
	if err != nil {
		hlog.Error(err)
		return
	}

	return
}
