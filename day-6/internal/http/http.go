package http

import (
	"day-6/internal/app/auth"
	"day-6/internal/app/user"
	"day-6/internal/factory"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewHttp(e *echo.Echo, f *factory.Factory) {
	e.Validator = &CustomValidator{validator: validator.New()}
	v1 := e.Group("/v1")

	user.NewHandler(f).Route(v1.Group("/users"))
	auth.NewHandler(f).Route(v1.Group("/auth"))
}
