package di

import (
	"clean-app/driver"
	"clean-app/gateway"
	"clean-app/rest/handler"
	"clean-app/usecase"
)

func InitTodoHandler() *handler.TodoHandler {
	db := driver.ProvidedDatabaseConnection()
	driver := driver.ProvideTodoDriver(db)
	gateway := gateway.ProvideTodoPort(driver)
	usecase := usecase.ProvideTodoUsecase(gateway)
	handler := handler.ProvideTodoHandler(usecase)

	return handler
}