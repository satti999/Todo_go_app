package service

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/satti999/todoapp/src/model"
	"github.com/satti999/todoapp/src/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (service *TodoService) Create(context *fiber.Ctx) error {

	todo := model.Todo{}

	if err := context.BodyParser(&todo); err != nil {

		return err

	}

	err := service.repo.Create(&todo)

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create Todo"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Todo has been added"})
	return nil

}

func (service *TodoService) GetAllTodos(context *fiber.Ctx) error {

	todos, err := service.repo.GetAll()

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create employee"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User has been added",
		"todos":   todos,
	})
	return nil
}

func (service *TodoService) UpdateTodo(context *fiber.Ctx) error {

	id := context.Params("id")

	var todo model.Todo

	if err := context.BodyParser(&todo); err != nil {

		return err

	}

	todoID, err := strconv.Atoi(id)

	if err != nil {

		// ... handle error

		panic(err)

	}
	ID := uint(todoID)

	err = service.repo.Update(ID, &todo)

	if err != nil {

		context.Status(http.StatusBadRequest).JSON(

			&fiber.Map{"message": "could not update todo"})

		return err

	}

	context.Status(http.StatusOK).JSON(&fiber.Map{

		"message": "todo has been updated",

		"todo": todo,
	})

	return nil

}

func (service *TodoService) DeleteTodo(context *fiber.Ctx) error {

	id := context.Params("id")

	todoID, err := strconv.Atoi(id)

	if err != nil {

		// ... handle error

		panic(err)
	}
	ID := uint(todoID)

	err = service.repo.Delete(ID)

	if err != nil {

		context.Status(http.StatusBadRequest).JSON(

			&fiber.Map{"message": "could not delete todo"})

		return err

	}

	context.Status(http.StatusOK).JSON(&fiber.Map{

		"message": "todo deleted successfully",
	})

	return nil

}

func (service *TodoService) GetTodo(context *fiber.Ctx) error {

	id := context.Params("id")

	todoID, err := strconv.Atoi(id)

	if err != nil {
		panic(err)
	}

	ID := uint(todoID)

	todo, err := service.repo.Get(ID)

	if err != nil {

		context.Status(http.StatusBadRequest).JSON(

			&fiber.Map{"message": "could not get todo"})

		return err

	}

	context.Status(http.StatusOK).JSON(&fiber.Map{

		"message": "todo fetch succesfully",

		"todo": todo,
	})

	return nil

}

func (service *TodoService) GetTodoByUserID(context *fiber.Ctx) error {

	ID := context.Params("userid")
	userID, err := strconv.Atoi(ID)

	if err != nil {
		panic(err)
	}

	id := uint(userID)
	userTods, err := service.repo.FindTodoByUserId(id)

	if err != nil {

		context.Status(http.StatusBadRequest).JSON(

			&fiber.Map{"message": "could not get todo"})

		return err

	}

	context.Status(http.StatusOK).JSON(&fiber.Map{

		"message": "todo fetch succesfully",

		"todo": userTods,
	})

	return nil

}
