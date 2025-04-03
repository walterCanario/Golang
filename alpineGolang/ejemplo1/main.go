package main

import (
	"encoding/json"
	"net/http"
)

type MenuItem struct {
	Text string `json:"text"`
	Link string `json:"link"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/base.html")
	})

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

	http.ListenAndServe(":8080", nil)
}
