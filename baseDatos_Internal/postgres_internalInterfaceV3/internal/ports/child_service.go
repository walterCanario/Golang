package ports

import (
	"postgres_internalInterface/internal/domain"
)

// ChildService define el contrato para operaciones relacionadas con hijos.
type ChildService interface {
	CreateChild(child *domain.Child) error
	GetChildByID(id int) (*domain.Child, error)
	UpdateChild(child *domain.Child) error
	DeleteChild(id int) error
	GetChildrenByUserID(userID int) ([]*domain.Child, error)
	GetSiblings(childID int) ([]*domain.Child, error)
	GetAllChildren() ([]*domain.Child, error)
	GetChildWithUserDetails(childID int) (*domain.ChildWithUserDetails, error) // Nuevo méto
}
