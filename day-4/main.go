package main

import (
	"day-4/config"
	m "day-4/middlewares"
	"day-4/routes"
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	config.InitDB()

	// Check .env file
	if err := godotenv.Load(); err != nil {
		panic(errors.New("Error loading .env file"))
	}
	jwt := os.Getenv("SECRET_JWT")
	if jwt == "" {
		panic("SECRET JWT is required")
	}

	e := routes.New()
	m.LogMiddleware(e)
	e.Logger.Fatal(e.Start("localhost:8000"))
}
