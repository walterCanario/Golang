package services

import (
	"hexagonal/internal/domain"
	"hexagonal/internal/port"
)

type JugadorService struct {
	repo port.JugadorRepository
}

func NewJugadorService(repo port.JugadorRepository) *JugadorService {
	return &JugadorService{repo: repo}
}

func (s *JugadorService) RegistrarJugador(j *domain.Jugador) error {
	return s.repo.Create(j)
}

func (s *JugadorService) ObtenerPorID(id int64) (*domain.Jugador, error) {
	return s.repo.GetByID(id)
}

func (s *JugadorService) ActualizarJugador(j *domain.Jugador) error {
	return s.repo.Update(j)
}

func (s *JugadorService) EliminarJugador(id int64) error {
	return s.repo.Delete(id)
}

func (s *JugadorService) ObtenerTodos() ([]domain.Jugador, error) {
	return s.repo.GetAll()
}
