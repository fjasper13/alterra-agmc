package controllers

import (
	"day-3/lib/database"
	"day-3/middlewares"
	"day-3/models"
	"errors"
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

	user, err := database.GetUserById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func CreateUserController(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return err
	}

	// Validate User Input
	if err := c.Validate(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Validate same email
	resEmail, err := database.GetUserByEmail(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if resEmail.ID != 0 {
		return echo.NewHTTPError(http.StatusBadRequest, error.Error(errors.New("Email already used")))
	}

	res, err := database.CreateUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    res,
	})
}

func UpdateUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// validate user id
	jwtUserId := middlewares.ExtractTokenUser(c)
	if id != jwtUserId {
		return echo.NewHTTPError(http.StatusBadRequest, error.Error(errors.New("Cannot Edit Different User")))
	}

	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	// Validate User Input
	if err := c.Validate(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := database.UpdateUser(id, &user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	res.ID = uint(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   res,
	})
}

func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	// validate user id
	jwtUserId := middlewares.ExtractTokenUser(c)
	if id != jwtUserId {
		return echo.NewHTTPError(http.StatusBadRequest, error.Error(errors.New("Cannot Delete Different User")))
	}

	user := models.User{}
	c.Bind(&user)

	res, err := database.DeleteUser(id, &user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	res.ID = uint(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   res,
	})
}

func LoginController(c echo.Context) error {

	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	// Validate User Input
	if err := c.Validate(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := database.Login(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res.Token, err = middlewares.GenerateToken(int(user.ID))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   res,
	})
}
