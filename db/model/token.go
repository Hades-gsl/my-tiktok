package model

type UserToken struct {
	Token    string `gorm:"not null;primaryKey"`
	Username string `gorm:"not null;unique;size: 32"`
	UserID   uint32 `gorm:"not null;index"`
}
