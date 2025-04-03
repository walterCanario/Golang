package main

import (
	"encoding/json"
	"net/http"
	"path/filepath"
)

type MenuItem struct {
	Text string `json:"text"`
	Link string `json:"link"`
}

func main() {
	// Servir archivos estáticos con el tipo MIME correcto
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Establecer el tipo MIME correcto para archivos CSS
		if filepath.Ext(r.URL.Path) == ".css" {
			w.Header().Set("Content-Type", "text/css")
		}
		fs.ServeHTTP(w, r)
	})))

	// Servir la plantilla HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/base.html")
	})

	// Endpoint para el menú lateral
	http.HandleFunc("/sidenav", func(w http.ResponseWriter, r *http.Request) {
		menu := r.URL.Query().Get("menu")
		var items []MenuItem

		switch menu {
		case "comparativas":
			items = []MenuItem{
				{Text: "Comparar Ventas", Link: "#"},
				{Text: "Comparar Usuarios", Link: "#"},
			}
		case "reportes":
			items = []MenuItem{
				{Text: "Reporte de Ventas", Link: "#"},
				{Text: "Reporte de Clientes", Link: "#"},
			}
		case "georeferencia":
			items = []MenuItem{
				{Text: "Mapa de Clientes", Link: "#"},
				{Text: "Mapa de Sucursales", Link: "#"},
			}
		default:
			items = []MenuItem{
				{Text: "Bienvenido", Link: "#"},
				{Text: "Selecciona una opción del menú superior", Link: "#"},
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(items)
	})

	// Iniciar el servidor
	http.ListenAndServe(":8080", nil)
}
