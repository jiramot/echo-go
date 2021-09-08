package mocks

import (
	"github.com/stretchr/testify/mock"
	"jiramot/echo-go/internal/core/domain"
)

type MockEchoRepository struct {
	mock.Mock
}

func (m *MockEchoRepository) EchoMessage(message string) (domain.Echo, error) {
	ret := m.Called(message)
	return ret.Get(0).(domain.Echo), ret.Error(1)
}
