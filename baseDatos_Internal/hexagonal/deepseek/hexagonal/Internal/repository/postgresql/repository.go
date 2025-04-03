// Internal/Repository/Postgresql/repository.go
package postgresql

import (
	"database/sql"
	"hexagonal/internal/domain"
	"hexagonal/internal/port"
)

type jugadorRepository struct {
	db *sql.DB
}


func NewJugadorRepository(db *sql.DB) port.JugadorRepository {
	return &jugadorRepository{db: db}
}

func (r *jugadorRepository) Create(j *domain.Jugador) error {
	_, err := r.db.Exec(
		"INSERT INTO jugadores(id, nombre, posicion) VALUES($1, $2, $3)",
		j.ID, j.Nombre, j.Posicion,
	)
	return err
}

func (r *jugadorRepository) GetByID(id string) (*domain.Jugador, error) {
	jugador := &domain.Jugador{}
	err := r.db.QueryRow(
		"SELECT id, nombre, posicion FROM jugadores WHERE id = $1",
		id,
	).Scan(&jugador.ID, &jugador.Nombre, &jugador.Posicion)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, port.ErrJugadorNoEncontrado
		}
		return nil, err
	}
	return jugador, nil
}

func (r *jugadorRepository) Update(j *domain.Jugador) error {
	result, err := r.db.Exec(
		"UPDATE jugadores SET nombre = $1, posicion = $2 WHERE id = $3",
		j.Nombre, j.Posicion, j.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return port.ErrJugadorNoEncontrado
	}
	return nil
}

func (r *jugadorRepository) Delete(id string) error {
	result, err := r.db.Exec(
		"DELETE FROM jugadores WHERE id = $1",
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return port.ErrJugadorNoEncontrado
	}
	return nil
}

func (r *jugadorRepository) GetAll() ([]*domain.Jugador, error) {
	rows, err := r.db.Query("SELECT id, nombre, posicion FROM jugadores")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jugadores []*domain.Jugador
	for rows.Next() {
		j := &domain.Jugador{}
		err := rows.Scan(&j.ID, &j.Nombre, &j.Posicion)
		if err != nil {
			return nil, err
		}
		jugadores = append(jugadores, j)
	}
	return jugadores, nil
}
