package service

import (
    "github.com/matiasvarela/errors"
    "jiramot/echo-go/internal/core/domain"
    "jiramot/echo-go/internal/core/ports"
    "jiramot/echo-go/pkg/apperrors"
)

type service struct {
    echoRepository ports.EchoPort
}

func NewEchoService(echoRepository ports.EchoPort) *service {
    return &service{
        echoRepository: echoRepository,
    }
}

func (svr *service) GetEchoMessage(message string) (domain.Echo, error) {
    echo, err := svr.echoRepository.EchoMessage(message)
    if err != nil {
        return domain.Echo{}, errors.New(apperrors.Internal, err, "message", "cause message")
    }
    return echo, nil
}