package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	UserId  uint `json:"user_id" gorm:"not null"`
	VideoId uint `json:"video_id" gorm:"not null"`
}

func (f *Favorite) TableName() string {
	return "favorites"
}
