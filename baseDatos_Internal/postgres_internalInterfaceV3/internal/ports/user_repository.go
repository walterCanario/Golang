package ports

import (
	"postgres_internalInterface/internal/domain"
)

// UserRepository define el contrato para operaciones relacionadas con usuarios.
type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByID(id int) (*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(id int) error
	GetAllUsers() ([]*domain.User, error)
}
