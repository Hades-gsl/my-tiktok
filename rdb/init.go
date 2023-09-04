package rdb

import (
	"tiktok/config"

	"github.com/go-redis/redis/v8"
)

var RedisDB *redis.Client

func init() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr: config.Redis_addr,
		DB:   0,
	})
}
