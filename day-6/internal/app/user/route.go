package user

import (
	"day-6/constants"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (h *handler) Route(g *echo.Group) {
	// User
	g.GET("", h.GetUsers, middleware.JWT([]byte(constants.SECRET_JWT)))
	g.GET("/:id", h.GetUserById, middleware.JWT([]byte(constants.SECRET_JWT)))
	g.POST("", h.CreateUser)
	g.PUT("/:id", h.UpdateUser, middleware.JWT([]byte(constants.SECRET_JWT)))
	g.DELETE("/:id", h.DeleteUser, middleware.JWT([]byte(constants.SECRET_JWT)))
}
