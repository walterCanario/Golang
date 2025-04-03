package main

import (
	"log"
	"net/http"

	"postgres_internalInterface/config"
	"postgres_internalInterface/internal/adapters/handlers"
	"postgres_internalInterface/internal/adapters/repository"
	"postgres_internalInterface/internal/service"

	_ "github.com/lib/pq"
)

func main() {
	// Database connection
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}
	defer db.Close()

	// Dependency Injection
	userRepo := repository.NewPostgresUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Router setup
	r := http.NewServeMux()
	r.HandleFunc("GET /users", userHandler.GetAllUsers)
	r.HandleFunc("POST /users", userHandler.CreateUser)

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
