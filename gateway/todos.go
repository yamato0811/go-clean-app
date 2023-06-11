package gateway

import (
	"clean-app/domain"
	"clean-app/driver"
)

type TodoGateway struct {
	todoDriver driver.TodoDriver
}
// driverの役目は、DBやAPIのような外部のデータソースからアプリケーションの世界のデータに変換すること

func (t TodoGateway)GetTodos() ([]domain.Todo, error) {
	result, err := t.todoDriver.GetTodos() // json型で返ってくる
	if err != nil {
		return nil, err
	}

	var todos []domain.Todo
	for _, v := range result {
		todos = append(todos, domain.Todo{
			Id: domain.TodoId{Value: v.Id},
			Title: domain.TodoTitle{Value: v.Title},
			Completed: domain.TodoCompleted{Value: v.Completed},
		})
	}
	return todos, nil
}