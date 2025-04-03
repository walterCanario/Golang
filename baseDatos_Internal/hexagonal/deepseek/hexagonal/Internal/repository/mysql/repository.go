// Internal/Repository/Mysql/repository.go
package mysql

import (
	"database/sql"
	"errors"
	"hexagonal/internal/domain"
	"hexagonal/internal/port"
)

type jugadorRepository struct {
	db *sql.DB
}

var _ port.JugadorRepository = (*jugadorRepository)(nil)

func NewJugadorRepository(db *sql.DB) port.JugadorRepository {
	return &jugadorRepository{db: db}
}

func (r *jugadorRepository) Create(j *domain.Jugador) error {
	_, err := r.db.Exec(
		"INSERT INTO jugadores(id, nombre, posicion) VALUES(?, ?, ?)",
		j.ID, j.Nombre, j.Posicion,
	)
	return err
}

func (r *jugadorRepository) GetByID(id string) (*domain.Jugador, error) {
	jugador := &domain.Jugador{}
	err := r.db.QueryRow(
		"SELECT id, nombre, posicion FROM jugadores WHERE id = ?",
		id,
	).Scan(&jugador.ID, &jugador.Nombre, &jugador.Posicion)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, port.ErrJugadorNoEncontrado
		}
		return nil, err
	}
	return jugador, nil
}

func (r *jugadorRepository) Update(j *domain.Jugador) error {
	result, err := r.db.Exec(
		"UPDATE jugadores SET nombre = ?, posicion = ? WHERE id = ?",
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
		"DELETE FROM jugadores WHERE id = ?",
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

	jugadores := make([]*domain.Jugador, 0)
	for rows.Next() {
		j := &domain.Jugador{}
		if err := rows.Scan(&j.ID, &j.Nombre, &j.Posicion); err != nil {
			return nil, err
		}
		jugadores = append(jugadores, j)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return jugadores, nil
}
