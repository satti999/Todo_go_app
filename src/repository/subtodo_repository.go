package repository

import (
	"github.com/satti999/todoapp/src/model"
)

type SubTodoRepository struct {
	repo Repository
}

func NewSubTodoRepository(repo Repository) *SubTodoRepository {
	return &SubTodoRepository{
		repo: repo,
	}
}

func (repo *SubTodoRepository) Create(subTodo *model.SubTodo) error {
	return repo.repo.DB.Create(subTodo).Error
}

func (repo *SubTodoRepository) Get(id uint) (*model.SubTodo, error) {
	var subTodo model.SubTodo
	err := repo.repo.DB.First(&subTodo, id).Error
	return &subTodo, err
}

func (repo *SubTodoRepository) Update(subTodo *model.SubTodo) error {
	return repo.repo.DB.Save(subTodo).Error
}

func (repo *SubTodoRepository) Delete(id uint) error {
	return repo.repo.DB.Delete(&model.SubTodo{}, id).Error
}

func (repo *SubTodoRepository) GetAll() ([]model.SubTodo, error) {
	var subTodos []model.SubTodo
	err := repo.repo.DB.Find(&subTodos).Error
	return subTodos, err
}

func (repo *SubTodoRepository) GetByTodoId(id uint) ([]model.SubTodo, error) {
	var subTodos []model.SubTodo
	err := repo.repo.DB.Where("todo_id = ?", id).Find(&subTodos).Error
	return subTodos, err
}

func (repo *SubTodoRepository) DeleteByTodoId(id uint) error {
	return repo.repo.DB.Where("todo_id = ?", id).Delete(&model.SubTodo{}).Error
}
