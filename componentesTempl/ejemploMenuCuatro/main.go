package main

import (
	"context"
	"log"
	"net/http"

	"ejemploMenuCuatro/components"
	"ejemploMenuCuatro/templates"
)

func main() {
	// Ruta principal (Dashboard)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templates.DashboardPage().Render(context.Background(), w)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	// Ruta para el contenido del Dashboard
	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		err := templates.DashboardMenu().Render(context.Background(), w)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	// Ruta para el contenido de Comparativas
	http.HandleFunc("/comparativas", func(w http.ResponseWriter, r *http.Request) {
		err := components.ComparativasContent().Render(context.Background(), w)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	// Ruta para el contenido de Reportes
	http.HandleFunc("/reportes", func(w http.ResponseWriter, r *http.Request) {
		err := components.ReportesContent().Render(context.Background(), w)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	// Ruta para el contenido de Georeferencia
	http.HandleFunc("/georeferencia", func(w http.ResponseWriter, r *http.Request) {
		err := components.GeoreferenciaContent().Render(context.Background(), w)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":8090", nil))
}
