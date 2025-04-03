package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ChartData struct {
	Labels []string `json:"labels"`
	Data   []int    `json:"data"`
}

func getChartData(w http.ResponseWriter, r *http.Request) {
	// Configurar headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Permite cualquier origen
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	chartType := r.URL.Query().Get("type")

	var response ChartData

	switch chartType {
	case "barras":
		response = ChartData{
			Labels: []string{"Ene", "Feb", "Mar", "Abr", "May", "Jun"},
			Data:   []int{500, 700, 1200, 800, 1500, 1100},
		}
	case "lineas":
		response = ChartData{
			Labels: []string{"Ene", "Feb", "Mar", "Abr", "May", "Jun"},
			Data:   []int{400, 800, 1100, 900, 1400, 1200},
		}
	case "pastel":
		response = ChartData{
			Labels: []string{"Producto A", "Producto B", "Producto C", "Producto D", "Producto E"},
			Data:   []int{1048, 735, 580, 484, 300},
		}
	default:
		http.Error(w, "Tipo de gráfico no encontrado", http.StatusBadRequest)
		return
	}
	fmt.Println("\nData:")
	for _, data := range response.Data {
		fmt.Println(data)
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/chart-data", getChartData)
	http.ListenAndServe(":8080", nil)
}
