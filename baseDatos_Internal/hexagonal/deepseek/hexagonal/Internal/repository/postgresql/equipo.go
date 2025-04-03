package postgresql

import (
	"database/sql"
	"hexagonal/internal/domain"
	"hexagonal/internal/port"
)

type equipoRepository struct {
	db *sql.DB
}

func NewEquipoRepository(db *sql.DB) port.EquipoRepository {
	return &equipoRepository{db: db}
}

func (r *equipoRepository) Create(e *domain.Equipo) error {
	_, err := r.db.Exec(
		"INSERT INTO equipos(id, nombre, ciudad, estadio) VALUES($1, $2, $3, $4)",
		e.ID, e.Nombre, e.Ciudad, e.Estadio,
	)
	return err
}

func (r *equipoRepository) GetByID(id string) (*domain.Equipo, error) {
	equipo := &domain.Equipo{}
	err := r.db.QueryRow(
		"SELECT id, nombre, ciudad, estadio FROM equipos WHERE id = $1",
		id,
	).Scan(&equipo.ID, &equipo.Nombre, &equipo.Ciudad, &equipo.Estadio)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, port.ErrEquipoNoEncontrado
		}
		return nil, err
	}
	return equipo, nil
}

func (r *equipoRepository) Update(e *domain.Equipo) error {
	result, err := r.db.Exec(
		"UPDATE equipos SET nombre = $1, ciudad = $2, estadio = $3 WHERE id = $4",
		e.Nombre, e.Ciudad, e.Estadio, e.ID,
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return port.ErrEquipoNoEncontrado
	}
	return nil
}

func (r *equipoRepository) Delete(id string) error {
	result, err := r.db.Exec(
		"DELETE FROM equipos WHERE id = $1",
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return port.ErrEquipoNoEncontrado
	}
	return nil
}

func (r *equipoRepository) GetAll() ([]*domain.Equipo, error) {
	rows, err := r.db.Query("SELECT id, nombre, ciudad, estadio FROM equipos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var equipos []*domain.Equipo
	for rows.Next() {
		e := &domain.Equipo{}
		err := rows.Scan(&e.ID, &e.Nombre, &e.Ciudad, &e.Estadio)
		if err != nil {
			return nil, err
		}
		equipos = append(equipos, e)
	}
	return equipos, nil
}
