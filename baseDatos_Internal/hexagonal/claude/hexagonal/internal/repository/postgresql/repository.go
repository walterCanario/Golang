// internal/repository/postgresql/repository.go
package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	"hexagonal/internal/domain"

	_ "github.com/lib/pq"
)

type PostgresqlRepository struct {
	db *sql.DB
}

func NewPostgresqlRepository(host, port, user, password, dbname string) (*PostgresqlRepository, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresqlRepository{db: db}, nil
}

func (r *PostgresqlRepository) Create(ctx context.Context, jugador domain.Jugador) (domain.Jugador, error) {
	query := `INSERT INTO jugadores 
		(nombre, apellido, fecha_nacimiento, nacionalidad, posicion, equipo, dorsal, altura, peso, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) 
		RETURNING id`

	err := r.db.QueryRowContext(
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
	).Scan(&jugador.ID)

	if err != nil {
		return domain.Jugador{}, err
	}

	return jugador, nil
}

func (r *PostgresqlRepository) GetByID(ctx context.Context, id int) (domain.Jugador, error) {
	query := `SELECT id, nombre, apellido, fecha_nacimiento, nacionalidad, posicion, equipo, dorsal, altura, peso, created_at, updated_at 
		FROM jugadores WHERE id = $1`

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

func (r *PostgresqlRepository) Update(ctx context.Context, jugador domain.Jugador) (domain.Jugador, error) {
	query := `UPDATE jugadores SET 
		nombre=$1, apellido=$2, fecha_nacimiento=$3, nacionalidad=$4, posicion=$5, 
		equipo=$6, dorsal=$7, altura=$8, peso=$9, updated_at=$10 
		WHERE id=$11`

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

func (r *PostgresqlRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM jugadores WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *PostgresqlRepository) GetAll(ctx context.Context) ([]domain.Jugador, error) {
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
