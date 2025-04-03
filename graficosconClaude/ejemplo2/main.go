package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"
)

// Estructura de filtros
type Filtros struct {
	Nacionalidad string `json:"nacionalidad"`
	Educacion    string `json:"educacion"`
	Sexo         string `json:"sexo"`
}

// Estructura para el gráfico
type Grafico struct {
	ID     string                   `json:"id"`
	Datos  []map[string]interface{} `json:"datos"`
	Filtro Filtros                  `json:"filtros"`
}

// Datos de ejemplo (simulando una base de datos)
var datosGlobales = []map[string]string{
	{"nacionalidad": "Mexico", "educacion": "Universitario", "sexo": "Masculino", "valor": "100"},
	{"nacionalidad": "Mexico", "educacion": "Universitario", "sexo": "Femenino", "valor": "150"},
	{"nacionalidad": "USA", "educacion": "Bachillerato", "sexo": "Masculino", "valor": "200"},
	{"nacionalidad": "USA", "educacion": "Bachillerato", "sexo": "Femenino", "valor": "180"},
	{"nacionalidad": "Canada", "educacion": "Maestria", "sexo": "Masculino", "valor": "250"},
	{"nacionalidad": "Canada", "educacion": "Maestria", "sexo": "Femenino", "valor": "220"},
}

// Función para filtrar datos
func filtrarDatos(filtros Filtros) []map[string]interface{} {
	var datosFiltrados []map[string]interface{}

	for _, dato := range datosGlobales {
		if (filtros.Nacionalidad == "" || dato["nacionalidad"] == filtros.Nacionalidad) &&
			(filtros.Educacion == "" || dato["educacion"] == filtros.Educacion) &&
			(filtros.Sexo == "" || dato["sexo"] == filtros.Sexo) {
			datosFiltrados = append(datosFiltrados, map[string]interface{}{
				"name":  dato["nacionalidad"],
				"value": dato["valor"],
			})
		}
	}

	return datosFiltrados
}

// Función para generar un gráfico específico
func generarGrafico(tipo string, filtros Filtros, resultChan chan<- Grafico, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simular procesamiento
	time.Sleep(time.Second * 1)

	// Filtrar datos
	datos := filtrarDatos(filtros)

	grafico := Grafico{
		ID:     tipo,
		Datos:  datos,
		Filtro: filtros,
	}

	resultChan <- grafico
}

// Función para generar todos los gráficos concurrentemente
func generarGraficos(filtros Filtros) []Grafico {
	tiposGraficos := []string{"barras", "lineas", "pie"}
	resultChan := make(chan Grafico, len(tiposGraficos))
	var wg sync.WaitGroup

	// Lanzar goroutines para cada tipo de gráfico
	for _, tipo := range tiposGraficos {
		wg.Add(1)
		go generarGrafico(tipo, filtros, resultChan, &wg)
	}

	// Goroutine para cerrar el canal cuando se complete el procesamiento
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Recolectar resultados
	var graficos []Grafico
	for grafico := range resultChan {
		graficos = append(graficos, grafico)
	}

	return graficos
}

// Manejador HTTP para los gráficos
func handleGraficos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificar filtros del request
	var filtros Filtros
	if err := json.NewDecoder(r.Body).Decode(&filtros); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generar gráficos concurrentemente
	graficos := generarGraficos(filtros)

	// Configurar headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Enviar respuesta
	json.NewEncoder(w).Encode(graficos)
}

func main() {
	// Manejar preflight CORS
	http.HandleFunc("/graficos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}
		handleGraficos(w, r)
	})

	// Servir archivo HTML
	http.Handle("/", http.FileServer(http.Dir(".")))

	log.Println("Servidor iniciado en :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
