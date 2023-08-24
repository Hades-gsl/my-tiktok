package main

import (
	"log"
	"net"
	"tiktok/config"
	user "tiktok/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err.Error())
	}

	addr, err := net.ResolveTCPAddr("tcp", config.UserServiceAddr)
	if err != nil {
		log.Fatal(err.Error())
	}

	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user"}),
		server.WithRegistry(r), server.WithServiceAddr(addr))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
