package controllers

import (
	"day-4/lib/database"
	"day-4/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetBooksController(c echo.Context) error {
	books, err := database.GetBooks()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  books,
	})
}

func GetBookByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if id < 1 {
		return c.JSON(http.StatusBadRequest, errors.New("invalid id"))
	}

	book, err := database.GetBookById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   book,
	})
}

func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if book.Page == 0 {
		return c.JSON(http.StatusBadRequest, error.Error(errors.New("Invalid Request")))
	}

	res, err := database.CreateBook(&book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new book",
		"book":    res,
	})
}

func UpdateBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	book := models.Book{}
	c.Bind(&book)

	res, err := database.UpdateBook(id, &book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res.ID = uint(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   res,
	})
}

func DeleteBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	book := models.Book{}
	c.Bind(&book)

	res, err := database.DeleteBook(id, &book)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res.ID = uint(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   res,
	})
}
