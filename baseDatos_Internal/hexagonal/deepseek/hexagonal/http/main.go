// Http/main.go
package main

import (
	"database/sql"
	"log"
	"net/http"

	servicios "hexagonal/Internal/servicios/jugador"
	"hexagonal/handler/jugador"
	"hexagonal/internal/port"
	reposito "hexagonal/internal/repository"
	postgresRepo "hexagonal/internal/repository/postgresql"

	"github.com/gorilla/mux"
)

func main() {
	// Configurar router
	router := mux.NewRouter()

	// Configurar conexión a DB
	var db *sql.DB
	var err error
	var repo port.JugadorRepository

	// switch os.Getenv("DB_DRIVER") {
	// case "postgres":
	// 	db, err = repository.NewPostgresDB()
	// 	repo = postgresRepo.NewJugadorRepository(db)
	// case "mysql":
	// 	db, err = repository.NewMySQLDB()
	// 	repo = mysqlRepo.NewJugadorRepository(db)
	// default:
	// 	log.Fatal("DB_DRIVER no configurado")
	// }

	db, err = reposito.NewPostgresDB()
	repo = postgresRepo.NewJugadorRepository(db)

	if err != nil {
		log.Fatal("Error conectando a DB:", err)
	}
	defer db.Close()

	// Inyectar dependencias

	service := servicios.NewJugadorService(repo)
	handler := jugador.NewHandler(service)

	// Rutas
	router.HandleFunc("/jugadores", handler.CrearJugador).Methods("POST")
	// Agregar demás rutas...

	// Iniciar servidor
	log.Println("Servidor iniciado en :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
