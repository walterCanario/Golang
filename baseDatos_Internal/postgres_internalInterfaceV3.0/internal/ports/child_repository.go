package ports

import (
	"postgres_internalInterface/internal/domain"
)

// ChildRepository define el contrato para operaciones relacionadas con hijos.
type ChildRepository interface {
	CreateChild(child *domain.Child) error
	GetChildByID(id int) (*domain.Child, error)
	UpdateChild(child *domain.Child) error
	DeleteChild(id int) error
	GetChildrenByUserID(userID int) ([]*domain.Child, error)
	GetSiblings(childID int) ([]*domain.Child, error)
	GetAllChildren() ([]*domain.Child, error) // Agregamos este método
}
