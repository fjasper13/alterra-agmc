package user

import (
	"context"
	"day-7/internal/factory"
	"day-7/internal/models"
	"day-7/internal/repository"
)

type service struct {
	UserRepo repository.User
}

type Service interface {
	GetUsers(ctx context.Context) (*[]models.User, error)
	GetUserById(ctx context.Context, id int) (*models.User, error)
	GetUserByEmail(ctx context.Context, user *models.User) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	UpdateUser(ctx context.Context, id int, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id int) error
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepo: f.UserRepo,
	}
}

func (s *service) GetUsers(ctx context.Context) (*[]models.User, error) {

	users, err := s.UserRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *service) GetUserById(ctx context.Context, id int) (*models.User, error) {

	user, err := s.UserRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) GetUserByEmail(ctx context.Context, user *models.User) (*models.User, error) {

	user, _ = s.UserRepo.FindByEmail(ctx, user)

	return user, nil
}

func (s *service) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {

	user, err := s.UserRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) UpdateUser(ctx context.Context, id int, user *models.User) (*models.User, error) {

	user, err := s.UserRepo.Update(ctx, id, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) DeleteUser(ctx context.Context, id int) error {

	err := s.UserRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
