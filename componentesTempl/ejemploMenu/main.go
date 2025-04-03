// main.go
package main

import (
	"context"
	"log"
	"net/http"

	"ejemploMenu/templates"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templates.DashboardPage().Render(context.Background(), w)
		if err != nil {
			log.Printf("Error rendering template: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
