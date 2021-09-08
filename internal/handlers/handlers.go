package handler

import (
    "github.com/labstack/echo/v4"
    "jiramot/echo-go/internal/core/ports"
    "net/http"
)

type EchoHttpHandler struct {
    echoUseCase ports.EchoUseCase
}

func NewEchoHttpHandler(echoUseCase ports.EchoUseCase) *EchoHttpHandler {
    return &EchoHttpHandler{
        echoUseCase: echoUseCase,
    }
}

func (hdl *EchoHttpHandler) GetEchoMessage(ctx echo.Context) error {
    msg := ctx.QueryParam("message")
    domain, err := hdl.echoUseCase.GetEchoMessage(msg)

    if err != nil {
        return ctx.String(http.StatusOK, "Not OK")
    }

    return ctx.JSON(http.StatusOK, domain)
}
