package controllers

import (
	"day-4/config"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	bookJSONValid = `{
		"title":"Dasar",
		"page":432,
		"author": "joohnn"
	}`

	bookJSONInvalid = `{
		"title":"Dasar",
		"page":"432",
		"author": "joohnn"
	}`

	bookIdValid   = 1
	bookIdInvalid = -44
	bookIdStr     = "Id"
)

func TestCreateBookControllerValid(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodPost, "/v1/books", strings.NewReader(bookJSONValid))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, CreateBookController(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		body := rec.Body.String()
		assert.Contains(t, body, "title")
		assert.Contains(t, body, "page")
		assert.Contains(t, body, "author")
	}
}

func TestCreateBookControllerInValid(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodPost, "/v1/books", strings.NewReader(bookJSONInvalid))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, CreateBookController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestGetBookControllerValid(t *testing.T) {
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v1/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, GetBooksController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookByIdControllerValid(t *testing.T) {
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v1/books/{id}", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bookIdValidStr := strconv.Itoa(bookIdValid)
	c.SetParamNames("id")
	c.SetParamValues(bookIdValidStr)

	// Assertions
	if assert.NoError(t, GetBookByIdController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookByIdControllerInvalid(t *testing.T) {
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v1/books/{id}", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bookIdInvalidValidStr := strconv.Itoa(bookIdInvalid)
	c.SetParamNames("id")
	c.SetParamValues(bookIdInvalidValidStr)

	// Assertions
	if assert.NoError(t, GetBookByIdController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestGetBookByIdControllerInvalidIdString(t *testing.T) {
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v1/books/{id}", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues(bookIdStr)

	// Assertions
	if assert.NoError(t, GetBookByIdController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateBookControllerValid(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodPut, "/v1/books/{id}", strings.NewReader(bookJSONValid))
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bookIdValidStr := strconv.Itoa(bookIdValid)
	c.SetParamNames("id")
	c.SetParamValues(bookIdValidStr)

	// Assertions
	if assert.NoError(t, UpdateBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateBookControllerInvalidJsonInput(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodPut, "/v1/books/{id}", strings.NewReader(bookJSONInvalid))
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bookIdValidStr := strconv.Itoa(bookIdValid)
	c.SetParamNames("id")
	c.SetParamValues(bookIdValidStr)

	// Assertions
	if assert.NoError(t, UpdateBookController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateBookControllerInvalidIdString(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodPut, "/v1/books/{id}", strings.NewReader(bookJSONValid))
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues(bookIdStr)

	// Assertions
	if assert.NoError(t, UpdateBookController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestDeleteBookControllerValid(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodPut, "/v1/books/{id}", nil)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bookIdValidStr := strconv.Itoa(bookIdValid)
	c.SetParamNames("id")
	c.SetParamValues(bookIdValidStr)

	// Assertions
	if assert.NoError(t, DeleteBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteBookControllerInvalidDifferentId(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodDelete, "/v1/books/{id}", nil)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	bookIdValidStr := strconv.Itoa(bookIdInvalid)
	c.SetParamNames("id")
	c.SetParamValues(bookIdValidStr)

	// Assertions
	if assert.NoError(t, DeleteBookController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteBookControllerInvalidIdString(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodDelete, "/v1/books/{id}", nil)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues(bookIdStr)

	// Assertions
	if assert.NoError(t, DeleteBookController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
