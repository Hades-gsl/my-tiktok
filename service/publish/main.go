package main

import (
	"log"
	"net"
	"tiktok/config"
	publish "tiktok/kitex_gen/publish/publishservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err.Error())
	}

	addr, err := net.ResolveTCPAddr("tcp", config.PublishServiceAddr)
	if err != nil {
		log.Fatal(err.Error())
	}

	svr := publish.NewServer(new(PublishServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "publish"}),
		server.WithRegistry(r), server.WithServiceAddr(addr),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
