package database

import (
	"dynamic-api/config"
	"dynamic-api/models"
)

func GetUsers() (*[]models.User, error) {
	var users *[]models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id int) (*models.User, error) {
	var users *models.User

	if err := config.DB.First(&users, id).Error; err != nil {
		return nil, err
	}
	return users, nil
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
