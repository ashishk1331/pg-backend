package models

import "gorm.io/gorm"

type UserInfo struct {
	Id       int     `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID   int     `gorm:"not null" json:"userId"` // Foreign key
	Username string  `gorm:"type:varchar(255);unique;not null" json:"userName"`
	Fullname string  `gorm:"type:varchar(255);not null" json:"fullName"`
	ImageUrl *string `gorm:"type:varchar(255);default:null" json:"imageUrl"`
	User     User    `gorm:"foreignKey:UserID;references:Id" json:"user"` // UserID is the foreign key that references Id in the User table
}

func MigrateUserInfo(DB *gorm.DB) {
	DB.AutoMigrate(&UserInfo{})
}
