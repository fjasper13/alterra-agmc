package usermongo

import (
	"day-7/internal/factory"
	"day-7/internal/models"
	"net/http"

	"github.com/labstack/echo"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.FactoryMongo) *handler {
	return &handler{
		service: NewService(f),
	}
}

func (h *handler) GetUsersMongo(c echo.Context) error {

	users, err := h.service.GetUsersMongo(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func (h *handler) CreateUserMongo(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate User Input
	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := h.service.CreateUserMongo(c.Request().Context(), &user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new user",
		"user":    res,
	})
}
