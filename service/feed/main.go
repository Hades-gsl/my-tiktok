package main

import (
	"log"
	"net"
	"tiktok/config"
	feed "tiktok/kitex_gen/feed/feedservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err.Error())
	}

	addr, err := net.ResolveTCPAddr("tcp", config.FeedServiceAddr)
	if err != nil {
		log.Fatal(err.Error())
	}

	svr := feed.NewServer(new(FeedServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "feed"}),
		server.WithRegistry(r), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Fatal(err.Error())
	}
}
