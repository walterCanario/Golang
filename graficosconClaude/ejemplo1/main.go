package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// trabaja con el archivo indexSinBase.html
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

// Upgrader para conexiones WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Canal para sincronizar escritura de WebSocket
var (
	socketMutex sync.Mutex
	socketChan  = make(chan Grafico, 10)
)

// Función para filtrar datos
func filtrarDatos(filtros Filtros) []map[string]interface{} {
	var datosFiltrados []map[string]interface{}

	for _, dato := range datosGlobales {
		// Aplicar filtros
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

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Establish WebSocket connection
	cliente, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer cliente.Close()

	// Create a context for cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Channel for sending graphics
	socketChan := make(chan Grafico, 10)

	// Goroutine to send graphics
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case grafico, ok := <-socketChan:
				if !ok {
					return
				}
				socketMutex.Lock()
				err := cliente.WriteJSON(grafico)
				socketMutex.Unlock()

				if err != nil {
					log.Printf("Error sending graphic %s: %v", grafico.ID, err)
					cancel()
					return
				}
			}
		}
	}()

	// Handle multiple requests
	for {
		// Receive filters
		var filtros Filtros
		err = cliente.ReadJSON(&filtros)
		if err != nil {
			log.Println(err)
			cancel()
			close(socketChan)
			return
		}

		// Generate graphics safely
		go func(f Filtros) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Recovered from panic: %v", r)
				}
			}()
			generarGraficos(f, socketChan)
		}(filtros)
	}
}

// Modified generarGraficos to accept channel as parameter
func generarGraficos(filtros Filtros, socketChan chan Grafico) {
	tiposGraficos := []string{"barras", "lineas", "pie"}

	for _, tipo := range tiposGraficos {
		// Simulate processing time
		time.Sleep(time.Second * 1)

		// Filter data
		datos := filtrarDatos(filtros)

		grafico := Grafico{
			ID:     tipo,
			Datos:  datos,
			Filtro: filtros,
		}

		// Send to channel
		select {
		case socketChan <- grafico:
		default:
			log.Println("Channel full, skipping graphic")
		}
	}
}

func main() {
	http.HandleFunc("/graficos", handleWebSocket)

	// Servir archivo HTML
	http.Handle("/", http.FileServer(http.Dir(".")))

	log.Println("Servidor iniciado en :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
