package services

import (
	"coding_challenges/microapi/models"
	"coding_challenges/microapi/repositories"
)

type UserService interface {
	GetUser(id string) (*models.User, error)
	CreateUser(user *models.User) error
}

type DefaultUserService struct {
	repo repositories.UserRepository
}

func NewDefaultUserService(repo repositories.UserRepository) *DefaultUserService {
	return &DefaultUserService{repo: repo}
}

func (s *DefaultUserService) GetUser(id string) (*models.User, error) {
	return s.repo.GetUser(id)
}

func (s *DefaultUserService) CreateUser(user *models.User) error {
	return s.repo.CreateUser(user)
}
