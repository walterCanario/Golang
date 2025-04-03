package service

import (
	"errors"
	"postgres_internalInterface/internal/domain"
	"postgres_internalInterface/internal/ports"
)

// UserService define la estructura del servicio.
type UserService struct {
	repo ports.UserRepository // Usamos la interfaz
}

// NewUserService crea un nuevo servicio.
func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Métodos definidos
// CreateUser valida y pasa los datos al repositorio.
func (s *UserService) CreateUser(user *domain.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("nombre y correo electrónico son obligatorios")
	}
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id int) (*domain.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *domain.User) error {
	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}

func (s *UserService) GetAllUsers() ([]*domain.User, error) {
	return s.repo.GetAllUsers()
}
