package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func loadApp(port string) {
	e := echo.New()

	loadDotEnv()
	loadMiddleware(e)
	loadRoutes(e)

	e.Logger.Fatal(e.Start(port))
}

func loadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
