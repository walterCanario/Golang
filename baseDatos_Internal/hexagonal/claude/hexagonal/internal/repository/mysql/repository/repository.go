// internal/repository/mysql/repository.go
package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"hexagonal/internal/domain"

	_ "github.com/go-sql-driver/mysql"
)

type MysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(host, port, user, password, dbname string) (*MysqlRepository, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, password, host, port, dbname)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &MysqlRepository{db: db}, nil
}

func (r *MysqlRepository) Create(ctx context.Context, jugador domain.Jugador) (domain.Jugador, error) {
	query := `INSERT INTO jugadores 
		(nombre, apellido, fecha_nacimiento, nacionalidad, posicion, equipo, dorsal, altura, peso, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(
		ctx,
		query,
		jugador.Nombre,
		jugador.Apellido,
		jugador.FechaNac,
		jugador.Nacionalidad,
		jugador.Posicion,
		jugador.Equipo,
		jugador.Dorsal,
		jugador.Altura,
		jugador.Peso,
		jugador.CreatedAt,
		jugador.UpdatedAt,
	)

	if err != nil {
		return domain.Jugador{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return domain.Jugador{}, err
	}

	jugador.ID = int(id)
	return jugador, nil
}

func (r *MysqlRepository) GetByID(ctx context.Context, id int) (domain.Jugador, error) {
	query := `SELECT id, nombre, apellido, fecha_nacimiento, nacionalidad, posicion, equipo, dorsal, altura, peso, created_at, updated_at 
		FROM jugadores WHERE id = ?`

	var jugador domain.Jugador
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&jugador.ID,
		&jugador.Nombre,
		&jugador.Apellido,
		&jugador.FechaNac,
		&jugador.Nacionalidad,
		&jugador.Posicion,
		&jugador.Equipo,
		&jugador.Dorsal,
		&jugador.Altura,
		&jugador.Peso,
		&jugador.CreatedAt,
		&jugador.UpdatedAt,
	)

	if err != nil {
		return domain.Jugador{}, err
	}

	return jugador, nil
}

func (r *MysqlRepository) Update(ctx context.Context, jugador domain.Jugador) (domain.Jugador, error) {
	query := `UPDATE jugadores SET 
		nombre=?, apellido=?, fecha_nacimiento=?, nacionalidad=?, posicion=?, 
		equipo=?, dorsal=?, altura=?, peso=?, updated_at=? 
		WHERE id=?`

	_, err := r.db.ExecContext(
		ctx,
		query,
		jugador.Nombre,
		jugador.Apellido,
		jugador.FechaNac,
		jugador.Nacionalidad,
		jugador.Posicion,
		jugador.Equipo,
		jugador.Dorsal,
		jugador.Altura,
		jugador.Peso,
		jugador.UpdatedAt,
		jugador.ID,
	)

	if err != nil {
		return domain.Jugador{}, err
	}

	return jugador, nil
}

func (r *MysqlRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM jugadores WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *MysqlRepository) GetAll(ctx context.Context) ([]domain.Jugador, error) {
	query := `SELECT id, nombre, apellido, fecha_nacimiento, nacionalidad, posicion, 
		equipo, dorsal, altura, peso, created_at, updated_at FROM jugadores`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jugadores []domain.Jugador
	for rows.Next() {
		var jugador domain.Jugador
		err := rows.Scan(
			&jugador.ID,
			&jugador.Nombre,
			&jugador.Apellido,
			&jugador.FechaNac,
			&jugador.Nacionalidad,
			&jugador.Posicion,
			&jugador.Equipo,
			&jugador.Dorsal,
			&jugador.Altura,
			&jugador.Peso,
			&jugador.CreatedAt,
			&jugador.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		jugadores = append(jugadores, jugador)
	}

	return jugadores, nil
}
