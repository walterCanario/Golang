package mysql

import (
	"database/sql"
	"hexagonal/internal/domain"
	"hexagonal/internal/port"
)

type MySQLJugadorRepo struct {
	db *sql.DB
}

func NewMySQLJugadorRepo(db *sql.DB) port.JugadorRepository {
	return &MySQLJugadorRepo{db: db}
}

func (r *MySQLJugadorRepo) Create(j *domain.Jugador) error {
	query := `INSERT INTO jugadores (nombre, equipo, edad, posicion) VALUES (?, ?, ?, ?)`
	result, err := r.db.Exec(query, j.Nombre, j.Equipo, j.Edad, j.Posicion)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	j.ID = id
	return nil
}

func (r *MySQLJugadorRepo) GetByID(id int64) (*domain.Jugador, error) {
	query := `SELECT id, nombre, equipo, edad, posicion FROM jugadores WHERE id = ?`
	j := &domain.Jugador{}
	err := r.db.QueryRow(query, id).Scan(&j.ID, &j.Nombre, &j.Equipo, &j.Edad, &j.Posicion)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (r *MySQLJugadorRepo) Update(j *domain.Jugador) error {
	query := `UPDATE jugadores SET nombre=?, equipo=?, edad=?, posicion=? WHERE id=?`
	_, err := r.db.Exec(query, j.Nombre, j.Equipo, j.Edad, j.Posicion, j.ID)
	return err
}

func (r *MySQLJugadorRepo) Delete(id int64) error {
	query := `DELETE FROM jugadores WHERE id=?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *MySQLJugadorRepo) GetAll() ([]domain.Jugador, error) {
	query := `SELECT id, nombre, equipo, edad, posicion FROM jugadores`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jugadores []domain.Jugador
	for rows.Next() {
		var j domain.Jugador
		if err := rows.Scan(&j.ID, &j.Nombre, &j.Equipo, &j.Edad, &j.Posicion); err != nil {
			return nil, err
		}
		jugadores = append(jugadores, j)
	}
	return jugadores, nil
}
