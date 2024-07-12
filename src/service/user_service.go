package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/satti999/todoapp/src/model"
	"github.com/satti999/todoapp/src/repository"

	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) CreateUser(context *fiber.Ctx) error {
	var User model.User // Use the Employee type from the entity package
	if err := context.BodyParser(&User); err != nil {
		return err
	}
	fmt.Println("user data", User.Name)
	// Call repository method to save employee data to the database
	err := u.repo.CreateUser(User)
	// Return success response
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create employee"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "User has been added"})
	return nil
}

func (u *UserService) GetAllUsers(context *fiber.Ctx) error {

	users, err := u.repo.GetAllUsers()

	if err != nil {

		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get employees"})

		return err

	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":   "All employees fetch succesfully",
		"employees": users,
	})

	return nil
}

func (u *UserService) GetUser(context *fiber.Ctx) error {

	id := context.Params("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		// ... handle error
		panic(err)
	}
	//fmt.Println("user id", employeeID)
	ID := uint(userID)
	user, err := u.repo.GetUser(ID)

	if err != nil {

		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get user"})

		return err

	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user fetch succesfully",
		"user":    user,
	})

	return nil

}

func (u *UserService) UpdateUser(context *fiber.Ctx) error {

	id := context.Params("id")

	userID, err := strconv.Atoi(id)

	if err != nil {

		// ... handle error
		panic(err)

	}

	var user model.User

	if err := context.BodyParser(&user); err != nil {

		return err

	}

	ID := uint(userID)

	err = u.repo.UpdateUser(&user, ID)

	if err != nil {

		context.Status(http.StatusBadRequest).JSON(

			&fiber.Map{"message": "could not update user"})

		return err

	}
	context.Status(http.StatusOK).JSON(&fiber.Map{

		"message": "user updated succesfully",
	})

	return nil

}

func (u *UserService) DeleteUser(context *fiber.Ctx) error {

	id := context.Params("id")

	userID, err := strconv.Atoi(id)

	if err != nil {

		// ... handle error
		panic(err)

	}

	ID := uint(userID)

	err = u.repo.DeleteUser(ID)

	if err != nil {

		context.Status(http.StatusBadRequest).JSON(

			&fiber.Map{"message": "could not delete user"})

		return err

	}

	context.Status(http.StatusOK).JSON(&fiber.Map{

		"message": "user deleted succesfully",
	})

	return nil

}

func (u *UserService) GetUserByName(context *fiber.Ctx) error {

	name := context.Params("name")

	user, err := u.repo.GetUserByName(name)

	if err != nil {

		context.Status(http.StatusBadRequest).JSON(

			&fiber.Map{"message": "could not get user"})

		return err

	}

	context.Status(http.StatusOK).JSON(&fiber.Map{

		"message": "user fetch succesfully",
		"user":    user,
	})

	return nil

}
