package main

import (
	"github.com/labstack/echo/v4"
	"github.com/stefandorresteijn/goilerplate/controllers"
)

func loadRoutes(e *echo.Echo) {
	// Load all handlers here
	helloWorldController := controllers.HelloWorldController{}

	// Define all routes here
	e.GET("/", helloWorldController.Hello)
}
