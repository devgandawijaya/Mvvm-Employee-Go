package service

import (
	"go-api/model"
	"go-api/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	return s.Repo.GetAllUsers()
}

func (s *UserService) CreateUser(user model.User) error {
	return s.Repo.CreateUser(user)
}

func (s *UserService) UpdateUser(id int, user model.User) error {
	return s.Repo.UpdateUser(id, user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.Repo.DeleteUser(id)
}
