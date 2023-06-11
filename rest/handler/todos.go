package handler

import (
	"clean-app/domain"
	"clean-app/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoUseCase usecase.TodoUseCase
}

func (h TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.todoUseCase.GetTodos()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	response := TodoResponse{
		Todos: todos,
	}
	c.JSON(http.StatusOK, response)
}

type TodoResponse struct {
	Todos []domain.Todo `json:"todos"`
}