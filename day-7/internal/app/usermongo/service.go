package usermongo

import (
	"context"
	"day-7/internal/factory"
	"day-7/internal/models"
	"day-7/internal/repository"
)

type service struct {
	UserRepo repository.UserMongo
}

type Service interface {
	GetUsersMongo(ctx context.Context) (*[]models.User, error)
	CreateUserMongo(ctx context.Context, user *models.User) (*models.User, error)
}

func NewService(f *factory.FactoryMongo) Service {
	return &service{
		UserRepo: f.UserRepo,
	}
}

func (s *service) GetUsersMongo(ctx context.Context) (*[]models.User, error) {

	users, err := s.UserRepo.FindAllMongo(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *service) CreateUserMongo(ctx context.Context, user *models.User) (*models.User, error) {

	user, err := s.UserRepo.CreateMongo(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
