package model

import (
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	UserID    uint32
	Title     string `json:"title" gorm:"size:32"`
	FileAddr  string `json:"play_url"`
	CoverAddr string `json:"cover_url"`
	// FavoriteCount uint   `json:"favorite_count" gorm:"default:0"`
	// CommentCount  uint   `json:"comment_count" gorm:"default:0"`
}

func (v *Video) TableName() string {
	return "videos"
}
