package repository

import (
	"context"
	"day-7/internal/middlewares"
	"day-7/internal/models"

	"gorm.io/gorm"
)

type Auth interface {
	Login(ctx context.Context, data *models.User) (*models.User, error)
}

type auth struct {
	Db *gorm.DB
}

func NewAuth(db *gorm.DB) *auth {
	return &auth{
		db,
	}
}

func (r *auth) Login(ctx context.Context, user *models.User) (*models.User, error) {
	var err error

	if err := r.Db.WithContext(ctx).Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.GenerateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	return user, nil
}
