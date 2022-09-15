package middlewares

import (
	"day-3/config"
	"day-3/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func BasicAuthDB(email, password string, c echo.Context) (bool, error) {

	var user *models.User

	if err := config.DB.Where("email = ? AND password = ?", email, password).First(user).Error; err != nil {
		return false, err
	}
	return true, nil
}

func ExtractTokenUser(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userID := int(claims["userId"].(float64))
		return userID
	}
	return 0
}
