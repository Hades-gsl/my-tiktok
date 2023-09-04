package main

import (
	"log"
	"net"
	"tiktok/config"
	favorite "tiktok/kitex_gen/favorite/favoriteservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err.Error())
	}

	addr, err := net.ResolveTCPAddr("tcp", config.FavoriteServiceAddr)
	if err != nil {
		log.Fatal(err.Error())
	}

	svr := favorite.NewServer(new(FavoriteServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "favorite"}),
		server.WithRegistry(r), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
