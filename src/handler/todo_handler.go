package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satti999/todoapp/src/service"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(service *service.TodoService) *TodoHandler {
	return &TodoHandler{
		service: service,
	}
}

func TodoHandle(app *fiber.App, todoHandler *TodoHandler) {

	app.Post("/todos", todoHandler.service.Create)
	app.Get("/todos", todoHandler.service.GetAllTodos)

	app.Put("/todos/:id", todoHandler.service.UpdateTodo)
	app.Delete("/todos/:id", todoHandler.service.DeleteTodo)
	app.Get("/todos/:userid/GetTodoByUserID", todoHandler.service.GetTodoByUserID)
	app.Get("/todos/:id/GetTodo", todoHandler.service.GetTodo)
}
