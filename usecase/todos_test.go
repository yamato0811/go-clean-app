package usecase

import (
	"clean-app/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetTodos(t *testing.T) {
	todoPort := new(MockTodoPort)
	usecase := TodoUseCase{todoPort}
	todoPort.On("GetTodos").Return([]domain.Todo{}, nil)
	actual, _ := usecase.GetTodos()
	assert.Equal(t, []domain.Todo{}, actual)
}