package database

import (
	"day-10/internal/models"

	"gorm.io/gorm"
)

func InitMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}
