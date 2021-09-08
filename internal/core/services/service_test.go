package service_test

import (
    "github.com/stretchr/testify/assert"
    "jiramot/echo-go/internal/core/domain"
    "jiramot/echo-go/internal/core/services"
    "jiramot/echo-go/internal/mocks"
    "testing"
)

func TestGetEchoMessage(t *testing.T) {
    mockEchoRepository := new(mocks.MockEchoRepository)
    mockEchoDomain := &domain.Echo{
        Message: "Hello",
    }
    mockEchoRepository.On("EchoMessage", "Hello").Return(*mockEchoDomain, nil)

    echoService := service.NewEchoService(mockEchoRepository)
    message, err := echoService.GetEchoMessage("Hello")
    assert.Nil(t, err)
    assert.Equal(t, "Hello", message.Message)
}
