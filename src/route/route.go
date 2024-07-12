package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/satti999/todoapp/src/handler"
	"github.com/satti999/todoapp/src/repository"
	"github.com/satti999/todoapp/src/service"
	"gorm.io/gorm"
)

func MainRouteHandler(db *gorm.DB, app *fiber.App) {
	repo := repository.NewRepository(db)
	userRepo := repository.NewUserRepository(repo)
	todoRepo := repository.NewTodoRepository(repo)
	userSer := service.NewUserService(userRepo)
	todoSer := service.NewTodoService(todoRepo)
	userHandler := handler.NewUserHandler(userSer)
	todoHandler := handler.NewTodoHandler(todoSer)
	//route handler function
	handler.UserHandle(app, userHandler)
	handler.TodoHandle(app, todoHandler)

}
