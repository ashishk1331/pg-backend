package repository

import (
	"errors"
	"fmt"
	"pg-backend/config"
	"pg-backend/models"

	"gorm.io/gorm"
)

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user by Email: %w", err)
	}
	return &user, nil
}
func GetUserById(id int) (*models.User, error) {
	var user models.User
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, fmt.Errorf("failed to get user by user ID: %w", err)
	}
	return &user, nil
}
