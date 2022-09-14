package routes

import (
	"dynamic-api/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")
	v1.GET("/users", controllers.GetUsersController)
	v1.GET("/users/:id", controllers.GetUserByIdController)
	v1.POST("/users", controllers.CreateUserController)
	v1.PUT("/users/:id", controllers.UpdateUserController)
	v1.DELETE("/users/:id", controllers.DeleteUserController)
	return e
}
