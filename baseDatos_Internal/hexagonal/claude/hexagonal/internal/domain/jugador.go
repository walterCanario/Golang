// internal/domain/jugador.go
package domain

import (
	"time"
)

// Jugador es la entidad de dominio principal
type Jugador struct {
	ID           int       `json:"id"`
	Nombre       string    `json:"nombre"`
	Apellido     string    `json:"apellido"`
	FechaNac     time.Time `json:"fecha_nacimiento"`
	Nacionalidad string    `json:"nacionalidad"`
	Posicion     string    `json:"posicion"`
	Equipo       string    `json:"equipo"`
	Dorsal       int       `json:"dorsal"`
	Altura       float64   `json:"altura"`
	Peso         float64   `json:"peso"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// JugadorDTO para transferencia de datos
type JugadorDTO struct {
	Nombre       string    `json:"nombre"`
	Apellido     string    `json:"apellido"`
	FechaNac     time.Time `json:"fecha_nacimiento"`
	Nacionalidad string    `json:"nacionalidad"`
	Posicion     string    `json:"posicion"`
	Equipo       string    `json:"equipo"`
	Dorsal       int       `json:"dorsal"`
	Altura       float64   `json:"altura"`
	Peso         float64   `json:"peso"`
}

// ToJugador convierte un DTO a entidad de dominio
func (dto JugadorDTO) ToJugador() Jugador {
	now := time.Now()
	return Jugador{
		Nombre:       dto.Nombre,
		Apellido:     dto.Apellido,
		FechaNac:     dto.FechaNac,
		Nacionalidad: dto.Nacionalidad,
		Posicion:     dto.Posicion,
		Equipo:       dto.Equipo,
		Dorsal:       dto.Dorsal,
		Altura:       dto.Altura,
		Peso:         dto.Peso,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}
