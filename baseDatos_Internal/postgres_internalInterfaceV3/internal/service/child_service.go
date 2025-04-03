package service

import (
	"errors"
	"postgres_internalInterface/internal/domain"
	"postgres_internalInterface/internal/ports"
)

// ChildService implementa el contrato ChildService.
type ChildService struct {
	childRepo ports.ChildRepository
	userRepo  ports.UserRepository // Inyectamos el repositorio de usuarios para el ultimo metodo
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

// GetChildWithUserDetails obtiene los datos de un hijo y los combina con los detalles del usuario padre.
func (s *ChildService) GetChildWithUserDetails(childID int) (*domain.ChildWithUserDetails, error) {
	// Obtener el hijo
	child, err := s.childRepo.GetChildByID(childID)
	if err != nil {
		return nil, err
	}

	// Obtener los detalles del usuario padre
	user, err := s.userRepo.GetUserByID(child.UserID)
	if err != nil {
		return nil, err
	}

	// Combinar los datos en un solo objeto
	return &domain.ChildWithUserDetails{
		Child: *child,
		User:  *user,
	}, nil
}
