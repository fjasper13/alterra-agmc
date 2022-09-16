package routes

import (
	"day-4/constants"
	"day-4/controllers"

	// m "day-4/middlewares"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	// mEcho "github.com/labstack/echo/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func New() *echo.Echo {
	e := echo.New()

	// Validator for input User
	e.Validator = &CustomValidator{validator: validator.New()}

	// Group for version 1
	v1 := e.Group("/v1")
	// Login User
	v1.POST("/login", controllers.LoginController)

	// Not Authenticated
	v1.GET("/books", controllers.GetBooksController)
	v1.GET("/books/:id", controllers.GetBookByIdController)
	v1.POST("/users", controllers.CreateUserController)

	// Authenticated
	// v1.Use(mEcho.BasicAuth(m.BasicAuthDB))
	v1.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// User
	v1.GET("/users", controllers.GetUsersController)
	v1.GET("/users/:id", controllers.GetUserByIdController)
	v1.PUT("/users/:id", controllers.UpdateUserController)
	v1.DELETE("/users/:id", controllers.DeleteUserController)

	// Book
	v1.POST("/books", controllers.CreateBookController)
	v1.PUT("/books/:id", controllers.UpdateBookController)
	v1.DELETE("/books/:id", controllers.DeleteBookController)
	return e
}
