package database

import (
	"day-6/internal/models"

	"gorm.io/gorm"
)

func InitMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}
