package ports

import (
    "jiramot/echo-go/internal/core/domain"
)

type EchoPort interface {
    EchoMessage(message string) (domain.Echo, error)
}

type EchoUseCase interface {
    GetEchoMessage(message string) (domain.Echo, error)
}