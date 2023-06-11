package usecase

import (
	"clean-app/domain"
	"clean-app/usecase/port"
)

type TodoUseCase struct {
	todoPort port.TodoPort
}

func (u TodoUseCase) GetTodos() ([]domain.Todo, error) {
	todos, err := u.todoPort.GetTodos()

	if err != nil {
		return nil, err
	}

	return todos, nil
}