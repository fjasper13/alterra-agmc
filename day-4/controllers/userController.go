package controllers

import (
	"day-4/lib/database"
	"day-4/middlewares"
	"day-4/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if id < 1 {
		return c.JSON(http.StatusBadRequest, errors.New("invalid id"))
	}

	user, err := database.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func CreateUserController(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate User Input
	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate same email
	resEmail, err := database.GetUserByEmail(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if resEmail.ID != 0 {
		return c.JSON(http.StatusBadRequest, error.Error(errors.New("Email already used")))
	}

	res, err := database.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new user",
		"user":    res,
	})
}

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// validate user id
	jwtUserId := middlewares.ExtractTokenUser(c)
	if id != jwtUserId {
		return c.JSON(http.StatusBadRequest, errors.New("Cannot Edit Different User"))
	}

	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate User Input
	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := database.UpdateUser(id, &user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res.ID = uint(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   res,
	})
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// validate user id
	jwtUserId := middlewares.ExtractTokenUser(c)
	if id != jwtUserId {
		return c.JSON(http.StatusBadRequest, error.Error(errors.New("Cannot Delete Different User")))
	}

	user := models.User{}
	c.Bind(&user)

	res, err := database.DeleteUser(id, &user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
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
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate User Input
	if err := c.Validate(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := database.Login(&user)
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
