package ports

import (
    domain "jiramot/echo-go/internal/core/domain"
)

type EchoRepository interface {
    Echo(message string) (domain.Echo, error)
}

type EchoService interface {
    Echo(message string) (domain.Echo, error)
}