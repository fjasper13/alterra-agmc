package usermongo

import (
	"day-7/constants"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.GetUsersMongo, middleware.JWT([]byte(constants.SECRET_JWT)))
	g.POST("", h.CreateUserMongo)
}
