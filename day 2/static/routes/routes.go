package routes

import (
	"static-api/controllers"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	v1 := e.Group("/v1")
	v1.GET("/books", controllers.GetBook)
	v1.GET("/books/:id", controllers.GetBookById)
	v1.POST("/books", controllers.CreateBook)
	v1.PUT("/books/:id", controllers.UpdateBook)
	v1.DELETE("/books/:id", controllers.DeleteBook)
	return e
}
