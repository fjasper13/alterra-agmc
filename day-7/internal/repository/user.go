package repository

import (
	"context"
	"day-7/internal/models"

	"gorm.io/gorm"
)

type User interface {
	FindAll(ctx context.Context) (*[]models.User, error)
	Create(ctx context.Context, data *models.User) (*models.User, error)
	FindByEmail(context.Context, *models.User) (*models.User, error)
	FindByID(ctx context.Context, ID int) (*models.User, error)
	Update(ctx context.Context, id int, user *models.User) (*models.User, error)
	Delete(ctx context.Context, ID int) error
}

type user struct {
	Db *gorm.DB
}

func NewUser(db *gorm.DB) *user {
	return &user{
		db,
	}
}

func (r *user) FindAll(ctx context.Context) (*[]models.User, error) {
	var users *[]models.User

	query := r.Db.WithContext(ctx).Model(&models.User{})
	err := query.Find(&users).Error

	return users, err
}

func (r *user) FindByID(ctx context.Context, id int) (*models.User, error) {
	var user *models.User
	err := r.Db.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *user) FindByEmail(ctx context.Context, data *models.User) (*models.User, error) {
	var user *models.User
	err := r.Db.WithContext(ctx).Model(&models.User{}).Where("email = ?", data.Email).First(&user).Error
	return user, err
}

func (r *user) Create(ctx context.Context, user *models.User) (*models.User, error) {
	if err := r.Db.WithContext(ctx).Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *user) Update(ctx context.Context, id int, user *models.User) (*models.User, error) {

	if err := r.Db.WithContext(ctx).Model(&user).Where("id = ?", id).Updates(models.User{Name: user.Name, Email: user.Email, Password: user.Password}).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *user) Delete(ctx context.Context, id int) error {
	if err := r.Db.WithContext(ctx).Where("id = ?", id).Delete(&models.User{}).Error; err != nil {
		return err
	}
	return nil
}
