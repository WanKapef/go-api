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
		return errors.New("nome e email são obrigatórios")
	}

	return s.repo.Create(user)
}

// Read
func (s *UserService) ListUsers(limit, offset, page int) ([]model.User, error) {
	// validação de parâmetros
	if page > 0 && offset > 0 {
		return nil, errors.New("use page ou offset, não ambos")
	}
	// valores padrão para paginação
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	// traduz page para offset
	if page > 0 {
		offset = (page - 1) * limit
	}
	if offset < 0 {
		offset = 0
	}

	users, err := s.repo.List(limit, offset)

	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, errors.New("nenhum usuário encontrado")
	}

	return users, nil
}

func (s *UserService) ListByID(id int64) (*model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Update
func (s *UserService) UpdateUser(user *model.User) error {
	if user.ID == 0 {
		return errors.New("ID do usuário é obrigatório")
	}
	if user.Name == "" || user.Email == "" {
		return errors.New("nome e email são obrigatórios")
	}

	return s.repo.Update(user)
}

// Delete
func (s *UserService) DeleteUser(userID int64) error {
	if userID == 0 {
		return errors.New("ID do usuário é obrigatório")
	}

	return s.repo.Delete(userID)
}
