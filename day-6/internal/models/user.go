package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
	Token    string `json:"token" form:"token"`
}
