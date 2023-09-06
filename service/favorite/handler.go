package main

import (
	"context"
	"strconv"
	"tiktok/config"
	"tiktok/db"
	"tiktok/db/model"
	favorite "tiktok/kitex_gen/favorite"
	"tiktok/kitex_gen/feed"
	"tiktok/kitex_gen/user"
	"tiktok/rdb"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"gorm.io/gorm"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// Action implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) Action(ctx context.Context, req *favorite.ActionRequest) (resp *favorite.ActionResponse, err error) {
	user_id := req.UserId
	video_id := req.VideoId
	action_type := req.ActionType
	switch action_type {
	case 1:
		{
			err = like(ctx, user_id, video_id)
			if err != nil {
				hlog.Error(err)
				err = kerrors.NewBizStatusError(config.SQLSaveErrorStatusCode, config.SQLSaveErrorStatusMsg)
			}
		}
	case 2:
		{
			err = dislike(ctx, user_id, video_id)
			if err != nil {
				hlog.Error(err)
				err = kerrors.NewBizStatusError(config.SQLDeleteErrorStatusCode, config.SQLDeleteErrorStatusMsg)
			}
		}
	default:
		{
			hlog.Error(config.UnknownFavoriteTypeStatusMsg)
			err = kerrors.NewBizStatusError(config.UnknownFavoriteTypeStatusCode, config.UnknownFavoriteTypeStatusMsg)
		}
	}
	if err != nil {
		return
	}

	resp = &favorite.ActionResponse{
		StatusCode: config.OKStatusCode,
		StatusMsg:  &config.OKStatusMsg,
	}
	return
}

func like(ctx context.Context, user_id, video_id int64) (err error) {
	err = db.Q.Favorite.WithContext(ctx).Create(&model.Favorite{
		UserId:  uint(user_id),
		VideoId: uint(video_id),
	})
	if err != nil {
		return
	}

	c, err := rdb.RedisDB.Incr(ctx, strconv.FormatInt(video_id, 10)+"_likes").Result()
	hlog.Infof("like: video_id : %v, count: %v", video_id, c)

	return
}

func dislike(ctx context.Context, user_id, video_id int64) (err error) {
	favorites := db.Q.Favorite
	result, err := favorites.WithContext(ctx).Where(favorites.UserId.Eq(uint(user_id)), favorites.VideoId.Eq(uint(video_id))).Delete()
	if err != nil {
		hlog.Error(err)
		return
	}
	hlog.Info(result.RowsAffected)

	c, err := rdb.RedisDB.Decr(ctx, strconv.FormatInt(video_id, 10)+"_likes").Result()
	hlog.Infof("dislike: video_id : %v, count: %v", video_id, c)

	return
}

// List implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) List(ctx context.Context, req *favorite.ListRequest) (resp *favorite.ListResponse, err error) {
	user_id := req.UserId

	find, err := db.Q.Favorite.WithContext(ctx).Where(db.Q.Favorite.UserId.Eq(uint(user_id))).Find()
	if err != nil {
		hlog.Error(err)
		err = kerrors.NewBizStatusError(config.SQLQueryErrorStatusCode, config.SQLQueryErrorStatusMsg)
		return
	}

	videos, err := convert(ctx, find, uint(user_id))
	if err != nil {
		hlog.Error(err)
		err = kerrors.NewBizStatusError(config.SQLQueryErrorStatusCode, config.SQLQueryErrorStatusMsg)
		return
	}

	resp = &favorite.ListResponse{
		StatusCode: config.OKStatusCode,
		StatusMsg:  &config.OKStatusMsg,
		VideoList:  videos,
	}
	return
}

func convert(ctx context.Context, favorites []*model.Favorite, user_id uint) (videos []*feed.Video, err error) {
	find, err := db.Q.User.WithContext(ctx).Where(db.Q.User.ID.Eq(user_id)).First()
	if err != nil {
		hlog.Error(err)
		return
	}

	tmp := int64(0)
	u := &user.User{
		Id:              int64(find.ID),
		Name:            find.UserName,
		FollowCount:     &tmp,
		FollowerCount:   &tmp,
		IsFollow:        false,
		Avatar:          &find.Avatar,
		BackgroundImage: &find.BackgroundImage,
		Signature:       &find.Signature,
		TotalFavorited:  &tmp,
		WorkCount:       &tmp,
		FavoriteCount:   &tmp,
	}

	video := db.Q.Video
	videos = make([]*feed.Video, len(favorites))
	var v *model.Video
	for i, f := range favorites {
		v, err = video.WithContext(ctx).Where(video.ID.Eq(f.VideoId)).First()
		if err != nil {
			hlog.Error(err)
			return
		}

		var f, c int64
		f, c, err = rdb.GetLikesAndCommentsCount(ctx, int64(v.ID))
		if err != nil {
			hlog.Error(err)
			return
		}

		var isFavorite bool
		_, err = db.Q.Favorite.WithContext(ctx).Where(db.Q.Favorite.UserId.Eq(uint(user_id)), db.Q.Favorite.VideoId.Eq(v.ID)).First()
		if err == nil {
			isFavorite = true
		} else if err == gorm.ErrRecordNotFound {
			err = nil
			isFavorite = false
		} else {
			hlog.Error(err)
			return
		}

		vdo := &feed.Video{
			Id:            int64(v.ID),
			Author:        u,
			PlayUrl:       v.FileAddr,
			CoverUrl:      v.CoverAddr,
			FavoriteCount: f,
			CommentCount:  c,
			IsFavorite:    isFavorite,
			Title:         v.Title,
		}

		videos[i] = vdo
	}

	return
}
