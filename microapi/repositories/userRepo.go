package repositories

import (
	"coding_challenges/microapi/models"
	"errors"
)

type UserRepository interface {
	GetUser(id string) (*models.User, error)
	CreateUser(user *models.User) error
}

type InMemoryUserRepository struct {
	users map[string]*models.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[string]*models.User),
	}
}

func (r *InMemoryUserRepository) GetUser(id string) (*models.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepository) CreateUser(user *models.User) error {
	r.users[user.ID] = user
	return nil
}
