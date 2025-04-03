package port

import "hexagonal/internal/domain"

type JugadorRepository interface {
	Create(jugador *domain.Jugador) error
	GetByID(id int64) (*domain.Jugador, error)
	Update(jugador *domain.Jugador) error
	Delete(id int64) error
	GetAll() ([]domain.Jugador, error)
}
