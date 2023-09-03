package db

import (
	"log"
	"tiktok/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB, err := gorm.Open(mysql.Open(config.DSN))
	if err != nil {
		log.Fatal(err.Error())
	}

	SetDefault(DB)
}
