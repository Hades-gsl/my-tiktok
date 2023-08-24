package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	DB, err := gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		log.Fatal(err.Error())
	}

	SetDefault(DB)

}
