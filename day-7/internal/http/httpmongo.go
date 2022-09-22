package http

import (
	"day-7/internal/app/usermongo"
	"day-7/internal/factory"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func NewHttpMongo(e *echo.Echo, f *factory.FactoryMongo) {
	e.Validator = &CustomValidator{validator: validator.New()}
	v1 := e.Group("/v1")

	usermongo.NewHandler(f).Route(v1.Group("/users"))
}
