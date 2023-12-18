package controllers

import (
	"github.com/labstack/echo/v4"
	helloworldView "github.com/stefandorresteijn/goilerplate/views/helloworld"
)

type HelloWorldController struct {
}

func (c HelloWorldController) Hello(ctx echo.Context) error {
	return render(ctx, helloworldView.Hello(helloworldView.HelloProps{Name: "World"}))
}
