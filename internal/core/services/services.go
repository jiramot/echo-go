package service

import (
    "github.com/matiasvarela/errors"
    "jiramot/echo-go/internal/core/domain"
    "jiramot/echo-go/internal/core/ports"
    "jiramot/echo-go/pkg/apperrors"
)

type service struct {
    echoPort ports.EchoPort
}

func NewEchoService(echoPort ports.EchoPort) *service {
    return &service{
        echoPort: echoPort,
    }
}

func (svr *service) GetEchoMessage(message string) (domain.Echo, error) {
    echo, err := svr.echoPort.EchoMessage(message)
    if err != nil {
        return domain.Echo{}, errors.New(apperrors.Internal, err, "message", "cause message")
    }
    return echo, nil
}
