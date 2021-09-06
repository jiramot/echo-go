package handler

import (
    "jiramot/echo-go/internal/core/ports"
    "net/http"
    "github.com/labstack/echo/v4"
)

type EchoHttpHandler struct {
    echoService ports.EchoService
}

func NewEchoHttpHandler(echoService ports.EchoService) *EchoHttpHandler {
    return &EchoHttpHandler{
        echoService: echoService,
    }
}

func (hdl *EchoHttpHandler) Echo(ctx echo.Context) error  {
    msg := ctx.QueryParam("message")
    domain, err := hdl.echoService.Echo(msg)

    if err != nil {
        return ctx.String(http.StatusOK, "Not OK")
    }

    return ctx.JSON(http.StatusOK, domain)
}