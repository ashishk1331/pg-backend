package repository

import (
	"errors"
	"fmt"
	"pg-backend/config"
	"pg-backend/models"

	"gorm.io/gorm"
)

func GetUserInfoByUserId(userId int) (*models.UserInfo, error) {
	var userInfo models.UserInfo
	if err := config.DB.Where("user_id = ?", userId).First(&userInfo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user by Email: %w", err)
	}
	return &userInfo, nil
}
