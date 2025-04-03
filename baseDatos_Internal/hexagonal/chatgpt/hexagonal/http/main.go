package main

import (
	"hexagonal/config"
	jugador "hexagonal/handler"
	"hexagonal/internal/repository/mysql"
	"hexagonal/internal/repository/postgresql"
	"hexagonal/internal/services"
	"log"
	"net/http"
	"os"
)

func main() {
	dbType := os.Getenv("DB_TYPE") // "mysql" o "postgres"

	db, err := config.ConnectDB(dbType)
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}
	defer db.Close()

	var repo services.JugadorService

	switch dbType {
	case "mysql":
		repo = *services.NewJugadorService(mysql.NewMySQLJugadorRepo(db))
	case "postgres":
		repo = *services.NewJugadorService(postgresql.NewPostgresJugadorRepo(db))
	default:
		log.Fatal("Base de datos no soportada")
	}

	handler := jugador.NewJugadorHandler(&repo)

	http.HandleFunc("/jugadores/create", handler.Create)
	http.HandleFunc("/jugadores/get", handler.GetByID)
	http.HandleFunc("/jugadores/getAll", handler.GetAll)
	http.HandleFunc("/jugadores/update", handler.Update)
	http.HandleFunc("/jugadores/delete", handler.Delete)

	log.Println("Servidor corriendo en :8080")
	http.ListenAndServe(":8080", nil)
}
