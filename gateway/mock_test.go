package gateway

import (
	"clean-app/driver"

	"github.com/stretchr/testify/mock"
)

type MockTodoDriver struct {
	mock.Mock
}

func (m *MockTodoDriver) GetTodos() ([]driver.Todo, error) {
	args := m.Called()
	return args.Get(0).([]driver.Todo), args.Error(1)
}