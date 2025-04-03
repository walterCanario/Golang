package servicios

import (
	"hexagonal/internal/domain"
	"hexagonal/internal/port"
)

type equipoService struct {
	repo port.EquipoRepository
}

var _ port.EquipoService = (*equipoService)(nil) // Verificación en tiempo de compilación

func NewEquipoService(repo port.EquipoRepository) port.EquipoService {
	return &equipoService{repo: repo}
}

func (s *equipoService) CrearEquipo(e *domain.Equipo) error {
	return s.repo.Create(e)
}

func (s *equipoService) ObtenerEquipo(id string) (*domain.Equipo, error) {
	return s.repo.GetByID(id)
}

func (s *equipoService) ActualizarEquipo(e *domain.Equipo) error {
	return s.repo.Update(e)
}

func (s *equipoService) EliminarEquipo(id string) error {
	return s.repo.Delete(id)
}

func (s *equipoService) ListarEquipos() ([]*domain.Equipo, error) {
	return s.repo.GetAll()
}
