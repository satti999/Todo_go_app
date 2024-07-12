package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satti999/todoapp/src/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func UserHandle(app *fiber.App, userHandler *UserHandler) {

	app.Post("/users", userHandler.service.CreateUser)
	app.Get("/users", userHandler.service.GetAllUsers)
	app.Get("/users/:id", userHandler.service.GetUser)
	app.Get("/users/:name", userHandler.service.GetUserByName)

	app.Put("/users/:id", userHandler.service.UpdateUser)
	app.Delete("/users/:id", userHandler.service.DeleteUser)

}
