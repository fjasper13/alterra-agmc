package main

import (
	"dynamic-api/config"
	"dynamic-api/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
