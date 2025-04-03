package postgresql

import (
	"database/sql"
	"hexagonal/internal/domain"
	"hexagonal/internal/port"
)

type PostgresJugadorRepo struct {
	db *sql.DB
}

func NewPostgresJugadorRepo(db *sql.DB) port.JugadorRepository {
	return &PostgresJugadorRepo{db: db}
}

func (r *PostgresJugadorRepo) Create(j *domain.Jugador) error {
	query := `INSERT INTO jugadores (nombre, equipo, edad, posicion) VALUES ($1, $2, $3, $4) RETURNING id`
	return r.db.QueryRow(query, j.Nombre, j.Equipo, j.Edad, j.Posicion).Scan(&j.ID)
}

func (r *PostgresJugadorRepo) GetByID(id int64) (*domain.Jugador, error) {
	query := `SELECT id, nombre, equipo, edad, posicion FROM jugadores WHERE id = $1`
	j := &domain.Jugador{}
	err := r.db.QueryRow(query, id).Scan(&j.ID, &j.Nombre, &j.Equipo, &j.Edad, &j.Posicion)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (r *PostgresJugadorRepo) Update(j *domain.Jugador) error {
	query := `UPDATE jugadores SET nombre=$1, equipo=$2, edad=$3, posicion=$4 WHERE id=$5`
	_, err := r.db.Exec(query, j.Nombre, j.Equipo, j.Edad, j.Posicion, j.ID)
	return err
}

func (r *PostgresJugadorRepo) Delete(id int64) error {
	query := `DELETE FROM jugadores WHERE id=$1`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *PostgresJugadorRepo) GetAll() ([]domain.Jugador, error) {
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
