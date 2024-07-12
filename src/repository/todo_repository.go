package repository

import (
	"fmt"

	"github.com/satti999/todoapp/src/model"
)

type TodoRepository struct {
	repo *Repository
}

func NewTodoRepository(repo *Repository) *TodoRepository {
	return &TodoRepository{
		repo: repo,
	}
}

func (repo *TodoRepository) Create(todo *model.Todo) error {

	err := repo.repo.DB.Create(todo).Error
	if err != nil {
		return err
	}
	user := model.User{}
	userID := todo.Createdby
	fmt.Println("user id in todo repo", userID)
	if err := repo.repo.DB.First(&user, userID).Error; err != nil {
		return err
	}
	user.Tods = append(user.Tods, *todo)
	// models.DB.Model(&blogPost).Updates(blogPost)
	repo.repo.DB.Model(&user).Updates(user)
	// if err := repo.repo.DB.Save(&user).Error; err != nil {
	// 	return err
	// }
	fmt.Println("user data in todo", user)
	return nil

}

func (repo *TodoRepository) Get(id uint) (*model.Todo, error) {
	var todo model.Todo
	err := repo.repo.DB.First(&todo, id).Error

	if err != nil {
		return nil, err

	}
	return &todo, err
}

func (repo *TodoRepository) Update(id uint, todo *model.Todo) error {
	Todo := model.Todo{}
	err := repo.repo.DB.Model(Todo).Where("id = ?", id).Updates(todo).Error
	if err != nil {
		return err
	}
	return err
	//return repo.repo.DB.Save(todo).Error
}

func (repo *TodoRepository) Delete(id uint) error {
	return repo.repo.DB.Delete(&model.Todo{}, id).Error
}

func (repo *TodoRepository) GetAll() ([]model.Todo, error) {
	var todos []model.Todo
	err := repo.repo.DB.Find(&todos).Error
	return todos, err
}

func (repo *TodoRepository) FindTodoByUserId(id uint) ([]model.Todo, error) {

	var todos []model.Todo

	err := repo.repo.DB.Where("createdby = ?", id).Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, err

}
