package main

import (
	"github.com/labstack/echo/v4"
	"jiramot/echo-go/internal/handlers"
	"jiramot/echo-go/internal/core/services"
	"jiramot/echo-go/internal/repositories"
)

func main() {
    echoRepository := repository.NewEchoRepository()
    echoService := service.New(echoRepository)
    echoHandler := handler.NewHttpHandler(echoService)

    e := echo.New()
    e.GET("/", echoHandler.Echo)
    e.Logger.Fatal(e.Start(":8080"))
}