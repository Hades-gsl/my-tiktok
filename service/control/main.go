package main

import (
	"tiktok/kitex_gen/favorite/favoriteservice"
	"tiktok/kitex_gen/feed/feedservice"
	"tiktok/kitex_gen/publish/publishservice"
	"tiktok/kitex_gen/user/userservice"
	"tiktok/service/control/mw"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	feedClient     feedservice.Client
	userClient     userservice.Client
	publishClinet  publishservice.Client
	favoriteClient favoriteservice.Client
)

func init() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		hlog.Fatal(err)
	}
	feedClient, err = feedservice.NewClient("feed", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}

	userClient, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}

	publishClinet, err = publishservice.NewClient("publish", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}

	favoriteClient, err = favoriteservice.NewClient("favorite", client.WithResolver(r))
	if err != nil {
		hlog.Fatal(err)
	}
}

func main() {
	h := server.Default()

	douyin := h.Group("/douyin")

	feed := douyin.Group("/feed")
	feed.Use(mw.AuthMiddleware())
	feed.GET("/", feedList)

	user := douyin.Group("/user")
	user.Use(mw.AuthMiddleware())
	user.POST("/register", register)
	user.POST("/login", login)
	user.GET("/", info)

	publish := douyin.Group("/publish")
	publish.Use(mw.JWTMiddleware.MiddlewareFunc())
	publish.POST("/action", publishAction)
	publish.GET("/list", publishList)

	favorite := douyin.Group("/favorite")
	favorite.Use(mw.JWTMiddleware.MiddlewareFunc())
	favorite.POST("/action", favoriteAction)
	favorite.GET("/list", favoriteList)

	h.Spin()
}
