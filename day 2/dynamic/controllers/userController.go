package controllers

import (
	"dynamic-api/lib/database"
	"dynamic-api/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := database.GetUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	res, err := database.CreateUser(&user)
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    res,
	})
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	c.Bind(&user)

	res, err := database.UpdateUser(id, &user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	res.ID = uint(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  res,
	})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	c.Bind(&user)

	res, err := database.DeleteUser(id, &user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	res.ID = uint(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  res,
	})
}
