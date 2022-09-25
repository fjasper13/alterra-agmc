package main

import (
	"day-10/database"
	"day-10/internal/factory"
	"day-10/internal/http"
	m "day-10/internal/middlewares"
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func main() {
	// Check .env file
	if err := godotenv.Load(); err != nil {
		panic(errors.New("Error loading .env file"))
	}
	jwt := os.Getenv("SECRET_JWT")
	if jwt == "" {
		panic("SECRET JWT is required")
	}

	database.InitDB()

	f := factory.NewFactory()
	// fMongo := factory.NewFactoryMongo()
	e := echo.New()
	http.NewHttp(e, f)
	m.LogMiddleware(e)

	port := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + port))
}
