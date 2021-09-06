package main

import (
	"github.com/labstack/echo/v4"
	handler "jiramot/echo-go/internal/handlers"
	service "jiramot/echo-go/internal/core/services"
	repository "jiramot/echo-go/internal/repositories"
)

func main() {
    echoRepository := repository.NewEchoRepository()
    echoService := service.New(echoRepository)
    echoHandler := handler.NewHttpHandler(echoService)

    e := echo.New()
    e.GET("/", echoHandler.Echo)
    e.Logger.Fatal(e.Start(":8080"))
}