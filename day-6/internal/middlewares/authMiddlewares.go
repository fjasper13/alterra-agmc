package middlewares

import (
	"day-6/database"
	"day-6/internal/models"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func BasicAuthDB(email, password string, c echo.Context) (bool, error) {

	var user *models.User

	if err := database.DB.Where("email = ? AND password = ?", email, password).First(user).Error; err != nil {
		return false, err
	}
	return true, nil
}

func ExtractTokenUser(e echo.Context) int {
	authorizationHeader := e.Request().Header.Get("Authorization")
	if !strings.Contains(authorizationHeader, "Bearer") {
		return 0
	}
	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return 0
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0
	}

	userId := claims["userId"].(float64)
	return int(userId)
}
