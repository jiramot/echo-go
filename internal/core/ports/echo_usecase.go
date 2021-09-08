package ports

import "jiramot/echo-go/internal/core/domain"

type EchoUseCase interface {
	GetEchoMessage(message string) (domain.Echo, error)
}
