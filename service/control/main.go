package main

import (
	"tiktok/kitex_gen/feed/feedservice"
	"tiktok/kitex_gen/user/userservice"
	"tiktok/service/control/mw"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	feedClient feedservice.Client
	userClient userservice.Client
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
}

func main() {
	h := server.Default()

	h.Use(mw.AuthMiddleware())

	douyin := h.Group("/douyin")

	douyin.GET("/feed", feedList)

	user := douyin.Group("/user")
	user.POST("/register", register)
	user.POST("/login", login)
	user.GET("/", info)

	h.Spin()
}
