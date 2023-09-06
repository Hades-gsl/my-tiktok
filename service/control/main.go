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
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
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

	var opts []client.Option
	opts = append(opts, client.WithResolver(r))
	opts = append(opts, client.WithTransportProtocol(transport.TTHeader))
	opts = append(opts, client.WithMetaHandler(transmeta.ClientTTHeaderHandler))

	feedClient, err = feedservice.NewClient("feed", opts...)
	if err != nil {
		hlog.Fatal(err)
	}

	userClient, err = userservice.NewClient("user", opts...)
	if err != nil {
		hlog.Fatal(err)
	}

	publishClinet, err = publishservice.NewClient("publish", opts...)
	if err != nil {
		hlog.Fatal(err)
	}

	favoriteClient, err = favoriteservice.NewClient("favorite", opts...)
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
