package main

import (
	"log"
	feed "tiktok/kitex_gen/feed/feedservice"
)

func main() {
	svr := feed.NewServer(new(FeedServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
