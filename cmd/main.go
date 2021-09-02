package main

import (
    "net/http"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

    e.GET("/", func(ctx echo.Context) error {
        return ctx.String(http.StatusOK, "Hello")
    })

    e.Logger.Fatal(e.Start(":8080"))
}