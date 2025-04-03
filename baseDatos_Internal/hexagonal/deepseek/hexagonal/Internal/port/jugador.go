package port

import (
	"errors"
	"hexagonal/internal/domain"
)

var (
	ErrJugadorNoEncontrado = errors.New("jugador no encontrado")
)

type JugadorRepository interface {
	Create(jugador *domain.Jugador) error
	GetByID(id string) (*domain.Jugador, error)
	Update(jugador *domain.Jugador) error
	Delete(id string) error
	GetAll() ([]*domain.Jugador, error)
}

type JugadorService interface {
	CrearJugador(jugador *domain.Jugador) error
	ObtenerJugador(id string) (*domain.Jugador, error)
	ActualizarJugador(jugador *domain.Jugador) error
	EliminarJugador(id string) error
	ListarJugadores() ([]*domain.Jugador, error)
}
