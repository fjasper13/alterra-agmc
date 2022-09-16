package database

import (
	"day-4/config"
	"day-4/middlewares"
	"day-4/models"
)

func Login(user *models.User) (*models.User, error) {
	var err error

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.GenerateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	return user, nil
}
