package rest

import (
	"clean-app/rest/handler"

	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")

	{
		SystemHandler := handler.NewSystemHandler()
		v1.GET("/system/ping", SystemHandler.Ping)
	}

	todos := v1.Group("/todos")
	{
		todos.GET("", handler.GetTodos)
	}

	return r
}