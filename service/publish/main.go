package main

import (
	"log"
	publish "tiktok/kitex_gen/publish/publishservice"
)

func main() {
	svr := publish.NewServer(new(PublishServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}