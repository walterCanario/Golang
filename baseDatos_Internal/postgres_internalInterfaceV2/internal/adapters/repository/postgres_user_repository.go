package repository

import (
	"database/sql"
	"postgres_internalInterface/internal/domain"
	"postgres_internalInterface/internal/ports"
)

// PostgresUserRepository implementa UserRepository para PostgreSQL.
type PostgresUserRepository struct {
	db *sql.DB
}

// NewPostgresUserRepository crea una nueva instancia del repositorio.
func NewPostgresUserRepository(db *sql.DB) ports.UserRepository {
	return &PostgresUserRepository{db: db}
}

// CreateUser inserta un nuevo usuario en PostgreSQL.
func (r *PostgresUserRepository) CreateUser(user *domain.User) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	return r.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
}

// GetUserByID obtiene un usuario por ID.
func (r *PostgresUserRepository) GetUserByID(id int) (*domain.User, error) {
	var user domain.User
	query := "SELECT id, name, email FROM users WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser actualiza los datos de un usuario.
func (r *PostgresUserRepository) UpdateUser(user *domain.User) error {
	query := "UPDATE users SET name = $1, email = $2 WHERE id = $3"
	_, err := r.db.Exec(query, user.Name, user.Email, user.ID)
	return err
}

// DeleteUser elimina un usuario por ID.
func (r *PostgresUserRepository) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

// GetAllUsers obtiene todos los usuarios.
func (r *PostgresUserRepository) GetAllUsers() ([]*domain.User, error) {
	query := "select regionid ,nombrerbd ,dsregion from public.avances_encuestas2"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
