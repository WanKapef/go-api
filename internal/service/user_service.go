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
