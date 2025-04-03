package repository

import (
	"database/sql"
	"postgres_internalInterface/internal/domain"
	"postgres_internalInterface/internal/ports"
)

// MySQLChildRepository implementa ChildRepository para MySQL.
type MySQLChildRepository struct {
	db *sql.DB
}

// NewMySQLChildRepository crea una nueva instancia del repositorio.
func NewMySQLChildRepository(db *sql.DB) ports.ChildRepository {
	return &MySQLChildRepository{db: db}
}

// CreateChild inserta un nuevo hijo en MySQL.
func (r *MySQLChildRepository) CreateChild(child *domain.Child) error {
	query := "INSERT INTO children (name, last_name, user_id) VALUES (?, ?, ?)"
	result, err := r.db.Exec(query, child.Name, child.LastName, child.UserID)
	if err != nil {
		return err
	}

	// Obtener el ID generado
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	child.ID = int(id)
	return nil
}

// GetChildByID obtiene un hijo por ID.
func (r *MySQLChildRepository) GetChildByID(id int) (*domain.Child, error) {
	var child domain.Child
	query := "SELECT id, name, last_name, user_id FROM children WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&child.ID, &child.Name, &child.LastName, &child.UserID)
	if err != nil {
		return nil, err
	}
	return &child, nil
}

// UpdateChild actualiza los datos de un hijo.
func (r *MySQLChildRepository) UpdateChild(child *domain.Child) error {
	query := "UPDATE children SET name = ?, last_name = ?, user_id = ? WHERE id = ?"
	_, err := r.db.Exec(query, child.Name, child.LastName, child.UserID, child.ID)
	return err
}

// DeleteChild elimina un hijo por ID.
func (r *MySQLChildRepository) DeleteChild(id int) error {
	query := "DELETE FROM children WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

// GetChildrenByUserID obtiene todos los hijos de un usuario.
func (r *MySQLChildRepository) GetChildrenByUserID(userID int) ([]*domain.Child, error) {
	query := "SELECT id, name, last_name, user_id FROM children WHERE user_id = ?"
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var children []*domain.Child
	for rows.Next() {
		var child domain.Child
		err := rows.Scan(&child.ID, &child.Name, &child.LastName, &child.UserID)
		if err != nil {
			return nil, err
		}
		children = append(children, &child)
	}
	return children, nil
}

// GetSiblings obtiene los hermanos de un hijo.
func (r *MySQLChildRepository) GetSiblings(childID int) ([]*domain.Child, error) {
	// Primero, obtenemos el user_id del hijo
	var userID int
	query := "SELECT user_id FROM children WHERE id = ?"
	err := r.db.QueryRow(query, childID).Scan(&userID)
	if err != nil {
		return nil, err
	}

	// Luego, obtenemos todos los hijos del mismo user_id, excluyendo al hijo actual
	return r.GetChildrenByUserID(userID)
}

// GetAllChildren obtiene todos los hijos.
func (r *MySQLChildRepository) GetAllChildren() ([]*domain.Child, error) {
	query := "SELECT id, name, last_name, user_id FROM children"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var children []*domain.Child
	for rows.Next() {
		var child domain.Child
		err := rows.Scan(&child.ID, &child.Name, &child.LastName, &child.UserID)
		if err != nil {
			return nil, err
		}
		children = append(children, &child)
	}
	return children, nil
}
