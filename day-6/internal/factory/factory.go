package factory

import (
	"day-6/database"
	"day-6/internal/repository"
)

type Factory struct {
	UserRepo repository.User
	AuthRepo repository.Auth
}

type FactoryMongo struct {
	UserRepo repository.UserMongo
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		UserRepo: repository.NewUser(db),
		AuthRepo: repository.NewAuth(db),
	}
}

func NewFactoryMongo() *FactoryMongo {
	client := database.GetConnectionMongo()
	return &FactoryMongo{
		UserRepo: repository.NewUserMongo(client),
	}
}
