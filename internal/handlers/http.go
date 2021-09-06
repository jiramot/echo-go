package handler

import (
    ports "jiramot/echo-go/internal/core/ports"
    "net/http"
    "github.com/labstack/echo/v4"
)

type HttpHandler struct {
    echoService ports.EchoService
}

func NewHttpHandler(echoService ports.EchoService) *HttpHandler {
    return &HttpHandler{
        echoService: echoService,
    }
}

func (hdl *HttpHandler) Echo(ctx echo.Context) error  {
//     domain, err := hdl.echoService.Echo("Hello")

//     if err != nil {
//         return ctx.String(http.StatusOK, "Not OK")
//     }

    return ctx.String(http.StatusOK, "Hello 1")
}