package main

import (
	"context"
	"tiktok/config"
	"tiktok/db"
	"tiktok/db/model"
	feed "tiktok/kitex_gen/feed"
	user "tiktok/kitex_gen/user"
	"tiktok/rdb"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"gorm.io/gorm"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// List implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) List(ctx context.Context, req *feed.ListRequest) (resp *feed.ListResponse, err error) {
	latesttime := req.LatestTime
	find, err := findVideos(ctx, *latesttime)
	if err != nil {
		hlog.Error(err)
		err = kerrors.NewBizStatusError(config.SQLQueryErrorStatusCode, config.SQLQueryErrorStatusMsg)
		return
	}

	if len(find) == 0 {
		hlog.Error(config.NoVideoStatusMsg)
		err = kerrors.NewBizStatusError(config.NoVideoStatusCode, config.NoVideoStatusMsg)
		return
	}

	nextTime := find[len(find)-1].CreatedAt.Unix()
	videos, err := convert(ctx, find, *req.UserId)
	if err != nil {
		hlog.Error(err)
		err = kerrors.NewBizStatusError(config.SQLQueryErrorStatusCode, config.SQLQueryErrorStatusMsg)
		return
	}

	resp = &feed.ListResponse{
		StatusCode: config.OKStatusCode,
		StatusMsg:  &config.OKStatusMsg,
		VideoList:  videos,
		NextTime:   &nextTime,
	}
	return
}

func findVideos(ctx context.Context, latesttime int64) ([]*model.Video, error) {
	video := db.Q.Video
	return video.WithContext(ctx).
		Where(video.CreatedAt.Lt(time.Unix(latesttime, 0))).
		Order(video.CreatedAt.Desc()).
		Limit(config.VideosCount).
		Find()
}

func convert(ctx context.Context, videos []*model.Video, actor_id int64) (res []*feed.Video, err error) {
	res = make([]*feed.Video, len(videos))
	for i, v := range videos {
		var author *user.User
		author, err = findUser(ctx, v.UserID)
		if err != nil {
			hlog.Error(err.Error())
		}

		var f, c int64
		f, c, err = rdb.GetLikesAndCommentsCount(ctx, int64(v.ID))
		if err != nil {
			hlog.Error(err)
			return
		}

		var isFavorite bool
		_, err = db.Q.Favorite.WithContext(ctx).Where(db.Q.Favorite.UserId.Eq(uint(actor_id)), db.Q.Favorite.VideoId.Eq(v.ID)).First()
		if err == nil {
			isFavorite = true
		} else if err == gorm.ErrRecordNotFound {
			err = nil
			isFavorite = false
		} else {
			hlog.Error(err)
			return
		}

		res[i] = &feed.Video{
			Id:            int64(v.ID),
			Author:        author,
			PlayUrl:       v.FileAddr,
			CoverUrl:      v.CoverAddr,
			FavoriteCount: f,
			CommentCount:  c,
			IsFavorite:    isFavorite,
			Title:         v.Title,
		}
	}
	return
}

func findUser(ctx context.Context, id uint32) (res *user.User, err error) {
	find, err := db.Q.User.WithContext(ctx).Where(db.Q.User.ID.Eq(uint(id))).First()
	if err != nil {
		return
	}
	res = &user.User{
		Id:              int64(find.ID),
		Name:            find.UserName,
		Avatar:          &find.Avatar,
		BackgroundImage: &find.BackgroundImage,
		Signature:       &find.Signature,
	}
	return
}
