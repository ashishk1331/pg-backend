package models

import "gorm.io/gorm"

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func MigrateUser(DB *gorm.DB) {
	DB.AutoMigrate(&User{})
}
