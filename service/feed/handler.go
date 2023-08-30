package main

import (
	"context"
	"tiktok/config"
	"tiktok/db"
	"tiktok/db/model"
	feed "tiktok/kitex_gen/feed"
	user "tiktok/kitex_gen/user"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// List implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) List(ctx context.Context, req *feed.ListRequest) (resp *feed.ListResponse, err error) {
	latesttime := req.LatestTime
	find, err := findVideos(ctx, *latesttime)
	if err != nil {
		resp = &feed.ListResponse{
			StatusCode: config.SQLQueryErrorStatusCode,
			StatusMsg:  &config.SQLQueryErrorStatusMsg,
		}
		return
	}

	if len(find) == 0 {
		resp = &feed.ListResponse{
			StatusCode: config.NoVideoStatusCode,
			StatusMsg:  &config.NoVideoStatusMsg,
		}
		return
	}

	nextTime := find[len(find)-1].CreatedAt.Unix()
	videos, err := convert(ctx, find)
	if err != nil {
		hlog.Error(err.Error())
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

func convert(ctx context.Context, videos []*model.Video) (res []*feed.Video, err error) {
	res = make([]*feed.Video, len(videos))
	for i, v := range videos {
		author, err := findUser(ctx, v.UserID)
		if err != nil {
			hlog.Error(err.Error())
		}

		res[i] = &feed.Video{
			Id:       int64(v.ID),
			Author:   author,
			PlayUrl:  v.FileAddr,
			CoverUrl: v.CoverAddr,
			Title:    v.Title,
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
