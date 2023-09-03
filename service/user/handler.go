package main

import (
	"context"
	"errors"
	"math/rand"
	"tiktok/config"
	"tiktok/db"
	"tiktok/db/model"
	user "tiktok/kitex_gen/user"
	"tiktok/service/control/mw"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	name := req.Username
	password := req.Password

	if name == "" || password == "" {
		resp = &user.RegisterResponse{
			StatusCode: config.NameOrPasswordEmptyStatusCode,
			StatusMsg:  &config.NameOrPasswordEmptyStatusMsg,
		}
		err = errors.New(config.NameOrPasswordEmptyStatusMsg)
		return
	}

	_, err = db.Q.User.WithContext(ctx).Where(db.Q.User.UserName.Eq(name)).First()
	if err != gorm.ErrRecordNotFound {
		if err != nil {
			hlog.Error(err.Error())
			resp = &user.RegisterResponse{
				StatusCode: config.SQLQueryErrorStatusCode,
				StatusMsg:  &config.SQLQueryErrorStatusMsg,
			}
		} else {
			resp = &user.RegisterResponse{
				StatusCode: config.NameExistStatusCode,
				StatusMsg:  &config.NameExistStatusMsg,
			}
			err = errors.New(config.NameExistStatusMsg)
		}
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		hlog.Error(err.Error())
		resp = &user.RegisterResponse{
			StatusCode: config.PasswordHashErrorStatusCode,
			StatusMsg:  &config.PasswordHashErrorStatusMsg,
		}
		return
	}

	err = db.Q.User.WithContext(ctx).Create(&model.User{
		UserName:        name,
		PassWord:        string(passwordHash),
		Avatar:          "https://api.multiavatar.com/" + randString(2, 6),
		BackgroundImage: "https://tuapi.eees.cc/api.php?category=dongman&type=302",
		Signature:       "TODO",
	})
	if err != nil {
		hlog.Error(err.Error())
		resp = &user.RegisterResponse{
			StatusCode: config.SQLSaveErrorStatusCode,
			StatusMsg:  &config.SQLSaveErrorStatusMsg,
		}
		return
	}

	u, err := db.Q.User.WithContext(ctx).Where(db.Q.User.UserName.Eq(name)).First()
	if err != nil {
		hlog.Error(err.Error())
		resp = &user.RegisterResponse{
			StatusCode: config.SQLQueryErrorStatusCode,
			StatusMsg:  &config.SQLQueryErrorStatusMsg,
		}
		return
	}

	token, _, err := mw.JWTMiddleware.TokenGenerator(u)
	if err != nil {
		hlog.Error(err)
		resp = &user.RegisterResponse{
			StatusCode: config.GenerateTokenErrorStatusCode,
			StatusMsg:  &config.GenerateTokenErrorStatusMsg,
		}
		return
	}

	resp = &user.RegisterResponse{
		StatusCode: config.OKStatusCode,
		StatusMsg:  &config.OKStatusMsg,
		UserId:     int64(u.ID),
		Token:      token,
	}
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	name := req.Username
	password := req.Password

	if name == "" || password == "" {
		resp = &user.LoginResponse{
			StatusCode: config.NameOrPasswordEmptyStatusCode,
			StatusMsg:  &config.NameOrPasswordEmptyStatusMsg,
		}
		err = errors.New(config.NameOrPasswordEmptyStatusMsg)
		return
	}

	u, err := db.Q.User.WithContext(ctx).Where(db.Q.User.UserName.Eq(name)).First()
	if err != nil {
		hlog.Error(err.Error())
		resp = &user.LoginResponse{
			StatusCode: config.UserNotFoundStatusCode,
			StatusMsg:  &config.UserNotFoundStatusMsg,
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PassWord), []byte(password))
	if err != nil {
		hlog.Error(err.Error())
		resp = &user.LoginResponse{
			StatusCode: config.PasswordWrongStatusCode,
			StatusMsg:  &config.PasswordWrongStatusMsg,
		}
		return
	}

	token, _, err := mw.JWTMiddleware.TokenGenerator(u)
	if err != nil {
		hlog.Error(err)
		resp = &user.LoginResponse{
			StatusCode: config.GenerateTokenErrorStatusCode,
			StatusMsg:  &config.GenerateTokenErrorStatusMsg,
		}
		return
	}

	resp = &user.LoginResponse{
		StatusCode: config.OKStatusCode,
		StatusMsg:  &config.OKStatusMsg,
		UserId:     int64(u.ID),
		Token:      token,
	}
	return
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, req *user.InfoRequest) (resp *user.InforResponse, err error) {
	id := req.UserId

	u, err := db.Q.User.Where(db.Q.User.ID.Eq(uint(id))).First()
	if err != nil {
		hlog.Error(err.Error())
		resp = &user.InforResponse{
			StatusCode: config.SQLQueryErrorStatusCode,
			StatusMsg:  &config.SQLQueryErrorStatusMsg,
		}
		return
	}

	tmp := int64(0)
	resp = &user.InforResponse{
		StatusCode: config.OKStatusCode,
		StatusMsg:  &config.OKStatusMsg,
		User: &user.User{
			Id:              int64(u.ID),
			Name:            u.UserName,
			FollowCount:     &tmp,
			FollowerCount:   &tmp,
			IsFollow:        false,
			Avatar:          &u.Avatar,
			BackgroundImage: &u.BackgroundImage,
			Signature:       &u.Signature,
			TotalFavorited:  &tmp,
			WorkCount:       &tmp,
			FavoriteCount:   &tmp,
		},
	}
	return
}

func randString(minLen, maxLen int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	length := seededRand.Intn(maxLen)
	if length < minLen {
		length = minLen
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
