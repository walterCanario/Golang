package port

import (
	"errors"
	"hexagonal/internal/domain"
)

var (
	ErrEquipoNoEncontrado = errors.New("Equipo no encontrado")
)

type EquipoRepository interface {
	Create(equipo *domain.Equipo) error
	GetByID(id string) (*domain.Equipo, error)
	Update(equipo *domain.Equipo) error
	Delete(id string) error
	GetAll() ([]*domain.Equipo, error)
}

type EquipoService interface {
	CrearEquipo(equipo *domain.Equipo) error
	ObtenerEquipo(id string) (*domain.Equipo, error)
	ActualizarEquipo(equipo *domain.Equipo) error
	EliminarEquipo(id string) error
	ListarEquipos() ([]*domain.Equipo, error)
}
