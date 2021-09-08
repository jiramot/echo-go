package ports

import "jiramot/echo-go/internal/core/domain"

// Service implement UseCase
type EchoUseCase interface {
    GetEchoMessage(message string) (domain.Echo, error)
}
