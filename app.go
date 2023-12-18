package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func loadApp(port string) {
	e := echo.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	loadMiddleware(e)
	loadRoutes(e)

	e.Logger.Fatal(e.Start(port))
}
