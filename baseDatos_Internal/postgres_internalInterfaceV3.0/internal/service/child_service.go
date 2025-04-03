package service

import (
	"errors"
	"postgres_internalInterface/internal/domain"
	"postgres_internalInterface/internal/ports"
)

// ChildService implementa el contrato ChildService.
type ChildService struct {
	childRepo ports.ChildRepository
}

// NewChildService crea un nuevo servicio.
func NewChildService(childRepo ports.ChildRepository) *ChildService {
	return &ChildService{childRepo: childRepo}
}

// CreateChild valida y pasa los datos al repositorio.
func (s *ChildService) CreateChild(child *domain.Child) error {
	if child.Name == "" || child.LastName == "" {
		return errors.New("nombre y apellido son obligatorios")
	}
	return s.childRepo.CreateChild(child)
}

func (s *ChildService) GetChildByID(id int) (*domain.Child, error) {
	return s.childRepo.GetChildByID(id)
}

func (s *ChildService) UpdateChild(child *domain.Child) error {
	return s.childRepo.UpdateChild(child)
}

func (s *ChildService) DeleteChild(id int) error {
	return s.childRepo.DeleteChild(id)
}

func (s *ChildService) GetChildrenByUserID(userID int) ([]*domain.Child, error) {
	return s.childRepo.GetChildrenByUserID(userID)
}

func (s *ChildService) GetSiblings(childID int) ([]*domain.Child, error) {
	return s.childRepo.GetSiblings(childID)
}

// GetAllChildren obtiene todos los hijos.
func (s *ChildService) GetAllChildren() ([]*domain.Child, error) {
	return s.childRepo.GetAllChildren()
}
