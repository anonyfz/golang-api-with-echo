package main

import (
	"go-echo-api/api"
	"go-echo-api/service"
	"go-echo-api/validation"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Validator = validation.NewCustomValidator()

	loginService := service.NewLoginService()
	loginAPI := api.NewLoginAPI(loginService)
	e.POST("/login", loginAPI.Login)

	e.Logger.Fatal(e.Start(":8080"))
}
