package gateway

import (
	"clean-app/domain"
	"clean-app/driver"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetTodos(t *testing.T) {
	mockDriver := new(MockTodoDriver)
	gateway := TodoGateway{mockDriver}
	todos := []driver.Todo{
		{Id: 1, Title: "title", Completed: false},
	}
	mockDriver.On("GetTodos").Return(todos, nil)
	actual, _ := gateway.GetTodos()
	expected := []domain.Todo{
		{Id: domain.TodoId{Value: 1}, Title: domain.TodoTitle{Value: "title"}, Completed: domain.TodoCompleted{false},},
	}
	assert.Equal(t, expected, actual)
}