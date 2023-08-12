package main

import (
	"log"
	chat "tiktok/kitex_gen/chat/chatservice"
)

func main() {
	svr := chat.NewServer(new(ChatServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
