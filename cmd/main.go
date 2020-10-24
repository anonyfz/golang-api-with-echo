package main

import (
	"go-echo-api/api"
	"go-echo-api/service"
	"go-echo-api/validation"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	enableDebugMode(e)
	enableValidation(e)
	enableRequestLog(e)
	enableRecovery(e)
	enableCORS(e)
	enableLoginAPI(e)
	e.Logger.Fatal(e.Start(":8080"))
}

func enableDebugMode(e *echo.Echo) {
	e.Debug = true
}

func enableValidation(e *echo.Echo) {
	e.Validator = validation.NewCustomValidator()
}

func enableCORS(e *echo.Echo) {
	e.Use(middleware.CORS())
}

func enableRequestLog(e *echo.Echo) {
	e.Use(middleware.Logger())
}

func enableRecovery(e *echo.Echo) {
	e.Use(middleware.Recover())
}

func enableLoginAPI(e *echo.Echo) {
	loginService := service.NewLoginService()
	loginAPI := api.NewLoginAPI(loginService)
	e.POST("/login", loginAPI.Login)
}
