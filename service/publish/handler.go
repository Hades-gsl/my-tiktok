package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"tiktok/config"
	"tiktok/db"
	"tiktok/db/model"
	"tiktok/kitex_gen/feed"
	publish "tiktok/kitex_gen/publish"
	"tiktok/kitex_gen/user"
	"tiktok/rdb"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
	"gorm.io/gorm"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// Action implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) Action(ctx context.Context, req *publish.ActionRequest) (resp *publish.ActionResponse, err error) {
	title := req.Title
	if title == "" {
		err = errors.New(config.TitleEmptyStatusMsg)
		hlog.Error(err)
		resp = &publish.ActionResponse{
			StatusCode: config.TitleEmptyStatusCode,
			StatusMsg:  &config.TitleEmptyStatusMsg,
		}
		return
	}

	data := req.Data
	fileType := http.DetectContentType(data)
	if fileType != "video/mp4" {
		err = errors.New(config.FileTypeErrorStatusMsg)
		hlog.Error(err)
		resp = &publish.ActionResponse{
			StatusCode: config.FileTypeErrorStatusCode,
			StatusMsg:  &config.FileTypeErrorStatusMsg,
		}
		return
	}

	name := uuid.NewString()
	err = os.WriteFile(config.VideoPath+name+".mp4", data, 0644)
	if err != nil {
		hlog.Error(err)
		resp = &publish.ActionResponse{
			StatusCode: config.FileSaveErrorStatusCode,
			StatusMsg:  &config.FileSaveErrorStatusMsg,
		}
		return
	}

	video := &model.Video{
		UserID: uint32(req.UserId),
		Title:  title,
		// FileAddr: "http://localhost:8080/videos/" + name + ".mp4",
		FileAddr: "http://10.0.2.2:8080/videos/" + name + ".mp4",
	}

	err = getCover(name, 1)
	if err != nil {
		hlog.Error(err)
	} else {
		// video.CoverAddr = "http://localhost:8080/covers/" + name + ".png"
		video.CoverAddr = "http://10.0.2.2:8080/covers/" + name + ".png"
	}

	err = db.Q.Video.WithContext(ctx).Save(video)
	if err != nil {
		hlog.Error(err)
		resp = &publish.ActionResponse{
			StatusCode: config.SQLSaveErrorStatusCode,
			StatusMsg:  &config.SQLSaveErrorStatusMsg,
		}
		return
	}

	v, err := db.Q.Video.WithContext(ctx).Where(db.Q.Video.FileAddr.Eq(video.FileAddr)).First()
	if err != nil {
		hlog.Error(err)
		resp = &publish.ActionResponse{
			StatusCode: config.SQLQueryErrorStatusCode,
			StatusMsg:  &config.SQLQueryErrorStatusMsg,
		}
		return
	}

	err = rdb.RedisDB.Set(ctx, strconv.FormatInt(int64(v.ID), 10)+"_likes", 0, 0).Err()
	if err != nil {
		hlog.Error(err)
		resp = &publish.ActionResponse{
			StatusCode: config.SQLSaveErrorStatusCode,
			StatusMsg:  &config.SQLSaveErrorStatusMsg,
		}
		return
	}

	err = rdb.RedisDB.Set(ctx, strconv.FormatInt(int64(v.ID), 10)+"_comments", 0, 0).Err()
	if err != nil {
		hlog.Error(err)
		resp = &publish.ActionResponse{
			StatusCode: config.SQLSaveErrorStatusCode,
			StatusMsg:  &config.SQLSaveErrorStatusMsg,
		}
		return
	}

	resp = &publish.ActionResponse{
		StatusCode: config.OKStatusCode,
		StatusMsg:  &config.OKStatusMsg,
	}
	return
}

func getCover(name string, frameNum int) (err error) {
	videoPath := config.VideoPath + name + ".mp4"
	coverPath := config.CoverPath + name + ".png"
	buf := bytes.NewBuffer(nil)
	err = ffmpeg_go.Input(videoPath).Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		hlog.Error(err)
		return
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		hlog.Error(err)
		return
	}

	err = imaging.Save(img, coverPath)
	if err != nil {
		hlog.Error(err)
		return
	}

	return
}

// List implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) List(ctx context.Context, req *publish.ListRequest) (resp *publish.ListResponse, err error) {
	user_id := req.UserId

	find, err := db.Q.Video.WithContext(ctx).Where(db.Q.Video.UserID.Eq(uint32(user_id))).Find()
	if err != nil {
		hlog.Error(err)
		resp = &publish.ListResponse{
			StatusCode: config.SQLQueryErrorStatusCode,
			StatusMsg:  &config.SQLQueryErrorStatusMsg,
		}
		return
	}

	videos, err := convert(ctx, find, user_id, req.ActorId)
	if err != nil {
		hlog.Error(err)
		resp = &publish.ListResponse{
			StatusCode: config.SQLQueryErrorStatusCode,
			StatusMsg:  &config.SQLQueryErrorStatusMsg,
		}
		return
	}

	resp = &publish.ListResponse{
		StatusCode: config.OKStatusCode,
		StatusMsg:  &config.OKStatusMsg,
		VideoList:  videos,
	}
	return
}

func convert(ctx context.Context, videos []*model.Video, user_id, actor_id int64) (res []*feed.Video, err error) {
	find, err := db.Q.User.WithContext(ctx).Where(db.Q.User.ID.Eq(uint(user_id))).First()
	if err != nil {
		return
	}

	author := &user.User{
		Id:              user_id,
		Name:            find.UserName,
		Avatar:          &find.Avatar,
		BackgroundImage: &find.BackgroundImage,
		Signature:       &find.Signature,
	}

	res = make([]*feed.Video, len(videos))
	for i, v := range videos {
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
