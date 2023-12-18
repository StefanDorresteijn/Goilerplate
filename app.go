package main

import "github.com/labstack/echo/v4"

func loadApp(port string) {
	e := echo.New()

	loadMiddleware(e)
	loadRoutes(e)

	e.Logger.Fatal(e.Start(port))
}
