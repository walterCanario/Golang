package main

import (
	"log"
	"net/http"

	"postgres_internalInterface/config"
	"postgres_internalInterface/internal/adapters/handlers/child"
	"postgres_internalInterface/internal/adapters/handlers/user"
	"postgres_internalInterface/internal/adapters/repository"
	"postgres_internalInterface/internal/service"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	// Conexión a PostgreSQL
	pgDB, err := config.InitDB()
	if err != nil {
		log.Fatalf("No se pudo conectar a PostgreSQL: %v", err)
	}
	defer pgDB.Close()

	// Conexión a MySQL
	mysqlDB, err := config.InitMySQLDB()
	if err != nil {
		log.Fatalf("No se pudo conectar a MySQL: %v", err)
	}
	defer mysqlDB.Close()

	// Dependency Injection
	userRepo := repository.NewPostgresUserRepository(pgDB)
	childRepo := repository.NewMySQLChildRepository(mysqlDB)

	userService := service.NewUserService(userRepo)
	childService := service.NewChildService(childRepo)

	// Crear instancias de los handlers
	userHandler := user.NewUserHandler(userService)
	childHandler := child.NewChildHandler(childService)

	// Configurar el enrutador
	r := http.NewServeMux()

	// Registrar las rutas de cada handler
	userHandler.RegisterRoutes(r)
	childHandler.RegisterRoutes(r)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
