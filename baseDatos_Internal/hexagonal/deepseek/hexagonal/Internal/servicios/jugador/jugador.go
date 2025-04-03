// Internal/servicios/jugador/jugador.go
package servicios

import (
	"hexagonal/internal/domain"
	"hexagonal/internal/port"
)

type jugadorService struct {
	repo port.JugadorRepository
}

var _ port.JugadorService = (*jugadorService)(nil) // Verificación en tiempo de compilación

func NewJugadorService(repo port.JugadorRepository) port.JugadorService {
	return &jugadorService{repo: repo}
}

func (s *jugadorService) CrearJugador(j *domain.Jugador) error {
	return s.repo.Create(j)
}

func (s *jugadorService) ObtenerJugador(id string) (*domain.Jugador, error) {
	return s.repo.GetByID(id)
}

func (s *jugadorService) ActualizarJugador(j *domain.Jugador) error {
	return s.repo.Update(j)
}

func (s *jugadorService) EliminarJugador(id string) error {
	return s.repo.Delete(id)
}

func (s *jugadorService) ListarJugadores() ([]*domain.Jugador, error) {
	return s.repo.GetAll()
}
