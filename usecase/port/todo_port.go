package port

import "clean-app/domain"

type TodoPort interface {
	GetTodos() ([]domain.Todo, error)
}