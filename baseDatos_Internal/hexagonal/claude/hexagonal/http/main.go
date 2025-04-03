// http/main.go
package main

import (
	"log"
	"net/http"
	"os"

	jugadorHandler "hexagonal/handler/jugador/jugador"
	port "hexagonal/hexagonal/internal/port/jugador"
	"hexagonal/internal/repository/mysql"
	"hexagonal/internal/repository/postgresql"
	jugadorService "hexagonal/internal/services/jugador"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	}

	// Configurar bases de datos
	postgresRepo, err := postgresql.NewPostgresqlRepository(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	mysqlRepo, err := mysql.NewMysqlRepository(
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DB"),
	)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	// Determinar qué repositorio usar según la configuración
	var jugadorRepo port.JugadorRepository
	dbType := os.Getenv("DB_TYPE")
	if dbType == "mysql" {
		jugadorRepo = mysqlRepo
	} else {
		// PostgreSQL por defecto
		jugadorRepo = postgresRepo
	}

	// Crear servicio con inyección de dependencia
	jugadorSvc := jugadorService.NewJugadorService(jugadorRepo)

	// Configurar manejadores HTTP
	router := mux.NewRouter()
	jugadorHandler := jugadorHandler.NewJugadorHandler(jugadorSvc)
	jugadorHandler.RegisterRoutes(router)

	// Iniciar servidor HTTP
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
