package service

import (
//     "github.com/matiasvarela/errors"
    ports "jiramot/echo-go/internal/core/ports"
    domain "jiramot/echo-go/internal/core/domain"
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
        return domain.Echo{}, nil
    }
    return echo, nil
}