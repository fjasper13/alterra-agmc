package auth

import (
	"github.com/labstack/echo"
)

func (h *handler) Route(g *echo.Group) {
	g.POST("/login", h.Login)
}
