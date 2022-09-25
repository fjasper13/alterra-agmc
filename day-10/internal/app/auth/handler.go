package auth

import (
	"day-10/internal/factory"
	"day-10/internal/middlewares"
	"day-10/internal/models"
	"net/http"

	"github.com/labstack/echo"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) Login(c echo.Context) error {

	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate User Input
	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := h.service.Login(c.Request().Context(), &user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res.Token, err = middlewares.GenerateToken(int(user.ID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   res,
	})
}
