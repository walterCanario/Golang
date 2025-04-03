// internal/services/jugador/jugador.go
package jugador

import (
	"context"
	"time"

	"hexagonal/internal/domain"
	"hexagonal/internal/port"
)

type Service interface {
	Create(ctx context.Context, jugadorDTO domain.JugadorDTO) (domain.Jugador, error)
	GetByID(ctx context.Context, id int) (domain.Jugador, error)
	Update(ctx context.Context, id int, jugadorDTO domain.JugadorDTO) (domain.Jugador, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]domain.Jugador, error)
}

type jugadorService struct {
	repository port.JugadorRepository
}

// NewJugadorService crea un nuevo servicio de jugador con inyección de dependencia
func NewJugadorService(repository port.JugadorRepository) Service {
	return &jugadorService{
		repository: repository,
	}
}

func (s *jugadorService) Create(ctx context.Context, jugadorDTO domain.JugadorDTO) (domain.Jugador, error) {
	jugador := jugadorDTO.ToJugador()
	return s.repository.Create(ctx, jugador)
}

func (s *jugadorService) GetByID(ctx context.Context, id int) (domain.Jugador, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *jugadorService) Update(ctx context.Context, id int, jugadorDTO domain.JugadorDTO) (domain.Jugador, error) {
	// Verificar si el jugador existe
	jugador, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return domain.Jugador{}, err
	}

	// Actualizar campos
	jugador.Nombre = jugadorDTO.Nombre
	jugador.Apellido = jugadorDTO.Apellido
	jugador.FechaNac = jugadorDTO.FechaNac
	jugador.Nacionalidad = jugadorDTO.Nacionalidad
	jugador.Posicion = jugadorDTO.Posicion
	jugador.Equipo = jugadorDTO.Equipo
	jugador.Dorsal = jugadorDTO.Dorsal
	jugador.Altura = jugadorDTO.Altura
	jugador.Peso = jugadorDTO.Peso
	jugador.UpdatedAt = time.Now()

	return s.repository.Update(ctx, jugador)
}

func (s *jugadorService) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}

func (s *jugadorService) GetAll(ctx context.Context) ([]domain.Jugador, error) {
	return s.repository.GetAll(ctx)
}
