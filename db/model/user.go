package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName        string `json:"name" gorm:"size:32"`
	PassWord        string `gorm:"size:32"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	FollowCount     uint   `json:"follow_count" gorm:"default:0"`
	FollowerCount   uint   `json:"follower_count" gorm:"default:0"`
	TotalFavorited  uint   `json:"total_favorited" gorm:"default:0"`
	WorkCount       uint   `json:"work_count" gorm:"default:0"`
	FavoriteCount   uint   `json:"favorite_count" gorm:"default:0"`
}

func (u *User) TableName() string {
	return "users"
}
