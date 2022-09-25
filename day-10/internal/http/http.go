package http

import (
	"day-10/internal/app/auth"
	"day-10/internal/app/user"
	"day-10/internal/factory"
	"net/http"

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
	e.GET("/", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})
	e.Validator = &CustomValidator{validator: validator.New()}
	v1 := e.Group("/v1")

	user.NewHandler(f).Route(v1.Group("/users"))
	auth.NewHandler(f).Route(v1.Group("/auth"))

	// mongo
	// usermongo.NewHandler(fMongo).Route(v1.Group("/users/mongo"))
}
