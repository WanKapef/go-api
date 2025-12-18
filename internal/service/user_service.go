package service

import (
	"errors"

	"github.com/WanKapef/go-api/internal/model"
	"github.com/WanKapef/go-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Create
func (s *UserService) CreateUser(user *model.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}

	return s.repo.Create(user)
}

// Read
func (s *UserService) ListUsers() ([]model.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("nenhum usuário encontrado")
	}

	return users, nil
}

// Update
func (s *UserService) UpdateUser(user *model.User) error {
	if user.ID == 0 {
		return errors.New("user ID is required")
	}
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}

	return s.repo.Update(user)
}

// Delete
func (s *UserService) DeleteUser(userID int64) error {
	if userID == 0 {
		return errors.New("user ID is required")
	}

	return s.repo.Delete(userID)
}
