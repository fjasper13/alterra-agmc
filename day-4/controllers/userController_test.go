package controllers

import (
	"day-4/config"
	"day-4/middlewares"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

var (
	userJSONValid = `{
		"name":"john",
		"email":"yy@mail.com",
		"password": "ggpass"
	}`

	userJSONInvalid = `{
		"name":"john",
		"email":"yyy",
		"password": "ggpass"
	}`

	userIdValid   = 24
	userIdInvalid = -24
	userIdStr     = "ID"
)

func TestCreateUserControllerValid(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(userJSONValid))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, CreateUserController(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		body := rec.Body.String()
		assert.Contains(t, body, "name")
		assert.Contains(t, body, "email")
		assert.Contains(t, body, "password")
	}
}

func TestCreateUserControllerInvalid(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(userJSONInvalid))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, CreateUserController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestLoginControllerValid(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodGet, "/v1/login", strings.NewReader(userJSONValid))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, LoginController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestLoginControllerInvalid(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(http.MethodGet, "/v1/login", strings.NewReader(userJSONInvalid))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, LoginController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestGetUsersControllerValid(t *testing.T) {
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUserByIdControllerValid(t *testing.T) {
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v1/users/{id}", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userIdValidStr := strconv.Itoa(userIdValid)
	c.SetParamNames("id")
	c.SetParamValues(userIdValidStr)

	// Assertions
	if assert.NoError(t, GetUserByIdController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUserByIdControllerInvalid(t *testing.T) {
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v1/users/{id}", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userIdInvalidValidStr := strconv.Itoa(userIdInvalid)
	c.SetParamNames("id")
	c.SetParamValues(userIdInvalidValidStr)

	// Assertions
	if assert.NoError(t, GetUserByIdController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestGetUserByIdControllerInvalidIdString(t *testing.T) {
	config.InitDB()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/v1/users/{id}", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetParamNames("id")
	c.SetParamValues(userIdStr)

	// Assertions
	if assert.NoError(t, GetUserByIdController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateUserControllerValid(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodPut, "/v1/users/{id}", strings.NewReader(userJSONValid))
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userToken, err := middlewares.GenerateToken(int(userIdValid))
	if err != nil {
		t.Fatal(err)
	}
	c.Request().Header.Add(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", userToken))

	userIdValidStr := strconv.Itoa(userIdValid)
	c.SetParamNames("id")
	c.SetParamValues(userIdValidStr)

	// Assertions
	if assert.NoError(t, UpdateUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUpdateUserControllerInvalidJsonInput(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodPut, "/v1/users/{id}", strings.NewReader(userJSONInvalid))
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userToken, err := middlewares.GenerateToken(int(userIdValid))
	if err != nil {
		t.Fatal(err)
	}
	c.Request().Header.Add(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", userToken))

	userIdValidStr := strconv.Itoa(userIdValid)
	c.SetParamNames("id")
	c.SetParamValues(userIdValidStr)

	// Assertions
	if assert.NoError(t, UpdateUserController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateUserControllerInvalidDifferentId(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodPut, "/v1/users/{id}", strings.NewReader(userJSONValid))
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userToken, err := middlewares.GenerateToken(int(userIdInvalid))
	if err != nil {
		t.Fatal(err)
	}
	c.Request().Header.Add(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", userToken))

	userIdValidStr := strconv.Itoa(userIdValid)
	c.SetParamNames("id")
	c.SetParamValues(userIdValidStr)

	// Assertions
	if assert.NoError(t, UpdateUserController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateUserControllerInvalidIdString(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodPut, "/v1/users/{id}", strings.NewReader(userJSONValid))
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userToken, err := middlewares.GenerateToken(int(userIdInvalid))
	if err != nil {
		t.Fatal(err)
	}
	c.Request().Header.Add(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", userToken))

	c.SetParamNames("id")
	c.SetParamValues(userIdStr)

	// Assertions
	if assert.NoError(t, UpdateUserController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestDeleteUserControllerValid(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodDelete, "/v1/users/{id}", nil)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userToken, err := middlewares.GenerateToken(int(userIdValid))
	if err != nil {
		t.Fatal(err)
	}
	c.Request().Header.Add(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", userToken))

	userIdValidStr := strconv.Itoa(userIdValid)
	c.SetParamNames("id")
	c.SetParamValues(userIdValidStr)

	// Assertions
	if assert.NoError(t, DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteUserControllerInvalidDifferentId(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodDelete, "/v1/users/{id}", nil)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userToken, err := middlewares.GenerateToken(int(userIdInvalid))
	if err != nil {
		t.Fatal(err)
	}
	c.Request().Header.Add(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", userToken))

	userIdValidStr := strconv.Itoa(userIdValid)
	c.SetParamNames("id")
	c.SetParamValues(userIdValidStr)

	// Assertions
	if assert.NoError(t, DeleteUserController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestDeleteUserControllerInvalidIdString(t *testing.T) {
	config.InitDB()
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	req := httptest.NewRequest(http.MethodDelete, "/v1/users/{id}", nil)
	req.Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userToken, err := middlewares.GenerateToken(int(userIdInvalid))
	if err != nil {
		t.Fatal(err)
	}
	c.Request().Header.Add(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", userToken))

	c.SetParamNames("id")
	c.SetParamValues(userIdStr)

	// Assertions
	if assert.NoError(t, DeleteUserController(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
