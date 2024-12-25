package models

import "gorm.io/gorm"

type User struct {
	Id        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Email     string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password  string `gorm:"type:text;not null" json:"password"`
	Role      string `gorm:"type:varchar(50);default:'user'" json:"role"`
	CreatedAt int    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int    `gorm:"autoUpdateTime" json:"updated_at"`
}

func MigrateUser(DB *gorm.DB) {
	DB.AutoMigrate(&User{})
}
