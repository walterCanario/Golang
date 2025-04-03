// internal/port/jugador.go
package port

import (
	"context"

	"hexagonal/internal/domain"
)

// JugadorRepository define el contrato para interactuar con el almacenamiento de jugadores
type JugadorRepository interface {
	Create(ctx context.Context, jugador domain.Jugador) (domain.Jugador, error)
	GetByID(ctx context.Context, id int) (domain.Jugador, error)
	Update(ctx context.Context, jugador domain.Jugador) (domain.Jugador, error)
	Delete(ctx context.Context, id int) error
	GetAll(ctx context.Context) ([]domain.Jugador, error)
}
