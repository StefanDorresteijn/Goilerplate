package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stefandorresteijn/goilerplate/controllers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	helloWorldController := controllers.HelloWorldController{}
	e.GET("/", helloWorldController.Hello)

	e.Logger.Fatal(e.Start(":3000"))
}
