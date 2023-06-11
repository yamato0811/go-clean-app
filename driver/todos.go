package driver

import "gorm.io/gorm"

type TodoDriver interface {
	GetTodos() ([]Todo, error)
}

type TodoDriverImpl struct {
	conn *gorm.DB
}

func (t TodoDriverImpl) GetTodos() ([]Todo, error) {
	todos := []Todo{}
	t.conn.Find(&todos)

	return todos, nil
}

func ProvideTodoDriver(conn *gorm.DB) TodoDriver {
	return TodoDriverImpl{conn: conn}
}

type Todo struct {
	Id int `gorm: "primary_key"`
	Title string `gorm: "type:size:255"`
	Completed bool `gorm: "type:bool"`
}