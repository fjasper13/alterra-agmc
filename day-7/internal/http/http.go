package http

import (
	"day-7/internal/app/auth"
	"day-7/internal/app/user"
	"day-7/internal/app/usermongo"
	"day-7/internal/factory"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewHttp(e *echo.Echo, f *factory.Factory, fMongo *factory.FactoryMongo) {
	e.Validator = &CustomValidator{validator: validator.New()}
	v1 := e.Group("/v1")

	user.NewHandler(f).Route(v1.Group("/users"))
	auth.NewHandler(f).Route(v1.Group("/auth"))

	// mongo
	usermongo.NewHandler(fMongo).Route(v1.Group("/users/mongo"))
}
