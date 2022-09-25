package user

import (
	"day-10/internal/factory"
	"day-10/internal/middlewares"
	"day-10/internal/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"

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

func (h *handler) GetUsers(c echo.Context) error {

	users, err := h.service.GetUsers(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func (h *handler) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if id < 1 {
		return c.JSON(http.StatusBadRequest, errors.New("invalid id"))
	}

	user, err := h.service.GetUserById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func (h *handler) CreateUser(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate User Input
	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate same email
	resEmail, err := h.service.GetUserByEmail(c.Request().Context(), &user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if resEmail.ID != 0 {
		return c.JSON(http.StatusBadRequest, error.Error(errors.New("Email already used")))
	}

	res, err := h.service.CreateUser(c.Request().Context(), &user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new user",
		"user":    res,
	})
}

func (h *handler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// validate user id
	jwtUserId, err := middlewares.ExtractTokenUser(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if id != jwtUserId {
		fmt.Println(jwtUserId)
		fmt.Println(id)
		return c.JSON(http.StatusBadRequest, error.Error(errors.New("Cannot Edit Different User")))
	}

	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate User Input
	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := h.service.UpdateUser(c.Request().Context(), id, &user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res.ID = uint(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   res,
	})
}

func (h *handler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// validate user id
	jwtUserId, err := middlewares.ExtractTokenUser(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if id != jwtUserId {
		return c.JSON(http.StatusBadRequest, error.Error(errors.New("Cannot Delete Different User")))
	}

	user := models.User{}
	c.Bind(&user)

	err = h.service.DeleteUser(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   "Success Delete User",
	})
}
