package repository

import (
	"github.com/satti999/todoapp/src/model"
)

type UserRepository struct {
	repo *Repository
}

func NewUserRepository(repo *Repository) *UserRepository {
	return &UserRepository{
		repo: repo,
	}
}

func (u *UserRepository) CreateUser(user model.User) error {
	return u.repo.DB.Create(&user).Error
}

func (u *UserRepository) GetUser(id uint) (*model.User, error) {
	var user model.User
	err := u.repo.DB.First(&user, id).Error
	return &user, err
}

func (u *UserRepository) UpdateUser(user *model.User, id uint) error {
	User := model.User{}
	err := u.repo.DB.Model(User).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return err
	}
	return err
}

func (u *UserRepository) DeleteUser(id uint) error {
	return u.repo.DB.Delete(&model.User{}, id).Error
}

func (u *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := u.repo.DB.Find(&users).Error
	return users, err
}

func (u *UserRepository) GetUserByName(name string) (*model.User, error) {
	var user model.User
	err := u.repo.DB.Where("name = ?", name).First(&user).Error
	return &user, err
}

func (u *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := u.repo.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (u *UserRepository) GetUserByPassword(password string) (*model.User, error) {
	var user model.User
	err := u.repo.DB.Where("password = ?", password).First(&user).Error
	return &user, err
}
