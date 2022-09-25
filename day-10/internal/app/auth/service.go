package auth

import (
	"context"
	"day-10/internal/factory"
	"day-10/internal/models"
	"day-10/internal/repository"
)

type service struct {
	AuthRepo repository.Auth
}

type Service interface {
	Login(ctx context.Context, user *models.User) (*models.User, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		AuthRepo: f.AuthRepo,
	}
}

func (s *service) Login(ctx context.Context, user *models.User) (*models.User, error) {
	users, err := s.AuthRepo.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	return users, nil
}
