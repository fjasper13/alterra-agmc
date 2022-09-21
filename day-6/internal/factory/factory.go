package factory

import (
	"day-6/database"
	"day-6/repository"
)

type Factory struct {
	UserRepo repository.User
	AuthRepo repository.Auth
}

func NewFactory() *Factory {
	db := database.GetConnection()
	return &Factory{
		UserRepo: repository.NewUser(db),
		AuthRepo: repository.NewAuth(db),
	}
}
