package services

import (
	"github.com/go-devops-demo/models"
	"github.com/go-devops-demo/repositories"
)

type UserService interface {
	GetUsers() ([]models.User, error)
	GetUserByID(id uint) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uint) error
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository}
}

func (s *userService) GetUsers() ([]models.User, error) {
	return s.repository.FindAll()
}

func (s *userService) GetUserByID(id uint) (models.User, error) {
	return s.repository.FindById(id)
}

func (s *userService) CreateUser(user models.User) (models.User, error) {
	return s.repository.Create(user)
}

func (s *userService) UpdateUser(user models.User) (models.User, error) {
	return s.repository.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	user, err := s.repository.FindById(id)
	if err != nil {
		return err
	}
	return s.repository.Delete(user)
}
