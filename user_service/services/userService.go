package services

import (
	"microservice_golang/user_service/models"
	"microservice_golang/user_service/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user models.User) error {
	return s.Repo.CreateUser(user)
}

func (s *UserService) GetAllUser() ([]models.User, error) {
	return s.Repo.GetAllUser()
}
