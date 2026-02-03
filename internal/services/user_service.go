package services

import (
	"errors"
	"go-crud-psql/internal/repositories"
)

type UserService interface {
	CreateUser(name, email string, age int) (*repositories.User, error)
	GetAllUsers() ([]repositories.User, error)
	GetUserByID(id uint) (*repositories.User, error)
	UpdateUser(id uint, name, email string, age int) (*repositories.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(name, email string, age int) (*repositories.User, error) {
	if name == "" || email == "" {
		return nil, errors.New("name and email are required")
	}

	user := &repositories.User{
		Name:  name,
		Email: email,
		Age:   age,
	}

	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetAllUsers() ([]repositories.User, error) {
	return s.repo.FindAll()
}

func (s *userService) GetUserByID(id uint) (*repositories.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) UpdateUser(id uint, name, email string, age int) (*repositories.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	if age > 0 {
		user.Age = age
	}

	err = s.repo.Update(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}