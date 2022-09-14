package controllers

import (
	"net/http"
	"static-api/models"
	"strconv"

	"github.com/labstack/echo"
)

func GetBook(c echo.Context) error {
	book := models.Book{Title: "Pemrograman", Page: 510, Author: "John"}
	return c.JSON(http.StatusOK, book)
}

func GetBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book := models.Book{Id: id, Title: "Pemrograman", Page: 510, Author: "John"}
	// Render Data - JSON Response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"book": book,
	})
}

func CreateBook(c echo.Context) error {
	// get data from value
	title := c.FormValue("title")
	pageStr := c.FormValue("page")
	author := c.FormValue("author")

	page, _ := strconv.Atoi(pageStr)
	var book models.Book
	book.Title = title
	book.Page = page
	book.Author = author

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create book",
		"book":     book,
	})
}

func UpdateBook(c echo.Context) error {
	// get data from value
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.FormValue("title")
	pageStr := c.FormValue("page")
	author := c.FormValue("author")

	page, _ := strconv.Atoi(pageStr)
	var book models.Book
	book.Id = id
	book.Title = title
	book.Page = page
	book.Author = author

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update book",
		"book":     book,
	})
}

func DeleteBook(c echo.Context) error {
	// get data from value
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.FormValue("title")
	pageStr := c.FormValue("page")
	author := c.FormValue("author")

	page, _ := strconv.Atoi(pageStr)
	var book models.Book
	book.Id = id
	book.Title = title
	book.Page = page
	book.Author = author

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete book",
		"book":     book,
	})
}
