package models

import "gorm.io/gorm"

type UserInfo struct {
	Id       int    `json:"id"`
	UserID   int    `json:"userId"`
	Username string `json:"userName"`
	Fullname string `json:"fullName"`
	ImageUrl string `json:"imageUrl"`
}

func MigrateUserInfo(DB *gorm.DB) {
	DB.AutoMigrate(&UserInfo{})
}
