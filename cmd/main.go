package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "jiramot/echo-go/internal/core/services"
    "jiramot/echo-go/internal/handlers"
    "jiramot/echo-go/internal/repositories"
)

func main() {
    echoRepository := repository.NewEchoRepository()
    echoService := service.NewEchoService(echoRepository)
    echoHandler := handler.NewEchoHttpHandler(echoService)

    e := echo.New()
    e.Use(middleware.Logger())
    e.GET("/", echoHandler.GetEchoMessage)
    e.Logger.Fatal(e.Start(":8080"))
}
