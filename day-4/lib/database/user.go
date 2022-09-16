package database

import (
	"day-4/config"
	"day-4/models"
	"errors"

	"gorm.io/gorm"
)

func GetUsers() (*[]models.User, error) {
	var users *[]models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id int) (*models.User, error) {
	var user *models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByEmail(user *models.User) (*models.User, error) {
	if err := config.DB.Where("email = ?", user.Email).First(user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			// handle the record not being found
			return nil, err
		}
	}
	return user, nil
}

func CreateUser(user *models.User) (*models.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(id int, user *models.User) (*models.User, error) {

	if err := config.DB.Model(&user).Where("id = ?", id).Updates(models.User{Name: user.Name, Email: user.Email, Password: user.Password}).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUser(id int, user *models.User) (*models.User, error) {

	if err := config.DB.Delete(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}
