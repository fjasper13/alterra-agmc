package main

import (
	"day-6/database"
	"day-6/internal/factory"
	"day-6/internal/http"
	m "day-6/internal/middlewares"
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
	fMongo := factory.NewFactoryMongo()
	e := echo.New()
	http.NewHttp(e, f, fMongo)
	m.LogMiddleware(e)

	e.Logger.Fatal(e.Start("localhost:8000"))
}
