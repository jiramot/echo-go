package service

import (
    "github.com/matiasvarela/errors"
    "jiramot/echo-go/pkg/apperrors"
    "jiramot/echo-go/internal/core/ports"
    "jiramot/echo-go/internal/core/domain"
)

type service struct {
    echoRepository ports.EchoRepository
}

func New(echoRepository ports.EchoRepository) *service {
    return &service{
        echoRepository: echoRepository,
    }
}

func (svr *service) Echo(message string) (domain.Echo, error) {
    echo, err := svr.echoRepository.Echo(message)
    if err != nil {
        return domain.Echo{}, errors.New(apperrors.Internal, err, "message", "cause message")
    }
    return echo, nil
}