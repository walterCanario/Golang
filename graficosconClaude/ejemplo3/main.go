package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/patrickmn/go-cache"
)

// Estructura de filtros
type Filtros struct {
	Nacionalidad string `json:"nacionalidad"`
	Educacion    string `json:"educacion"`
	Sexo         string `json:"sexo"`
}

// Estructura para el gráfico
type Grafico struct {
	ID    string         `json:"id"`
	Datos []GraficoDatos `json:"datos"`
}

// Estructura para los datos del gráfico
type GraficoDatos struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// Variables globales
var (
	nacionalidad sql.NullString
	db           *sql.DB
	c            *cache.Cache
)

// Configurar la conexión a la base de datos
func initDB() error {
	var err error
	connStr := "postgres://postgres:Sead_2023%23@192.168.8.2:5432/encvulne"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("No se puede conectar a la base de datos: %v", err)
	}
	log.Println("Conexión a la base de datos exitosa")
	return err
}

// Inicializar la conexión a la base de datos
// func initDB() error {
// 	var err error
// 	connStr := "postgres://postgres:contraseña@192.168.8.2:5432/nombre_db?sslmode=disable"
// 	db, err = sql.Open("postgres", connStr)
// 	if err != nil {
// 		return err
// 	}

// 	if err = db.Ping(); err != nil {
// 		return err
// 	}

// 	// Crear tabla si no existe
// 	// _, err = db.Exec(schema)
// 	return err
// }

func buildQuery(tipo string, filtros Filtros) (string, []interface{}) {
	var args []interface{}
	var conditions []string
	paramCount := 1

	// Base de la consulta, independientemente del tipo de gráfico
	baseQuery := "SELECT %s FROM comparativas WHERE 1=1"

	// Selección dinámica de métricas según el tipo de gráfico
	var selectFields string
	switch tipo {
	case "barras":
		selectFields = "nacionalidad as name, COUNT(*) as value, AVG(idencuesta) as salario_promedio"
	case "lineas":
		selectFields = "nacionalidad as name, COUNT(*) as value, AVG(idencuesta) as edad_promedio"
	case "pie":
		selectFields = "sexo as name, COUNT(*) as value, SUM(idencuesta) as antiguedad_total"
	default:
		return "", nil // Si el tipo no es válido, retorna vacío
	}

	// Aplicar filtros dinámicos
	if filtros.Nacionalidad != "" {
		conditions = append(conditions, fmt.Sprintf("nacionalidad = $%d", paramCount))
		args = append(args, filtros.Nacionalidad)
		paramCount++
	}
	if filtros.Educacion != "" {
		conditions = append(conditions, fmt.Sprintf("tipodependencia = $%d", paramCount))
		args = append(args, filtros.Educacion)
		paramCount++
	}
	if filtros.Sexo != "" {
		conditions = append(conditions, fmt.Sprintf("sexo = $%d", paramCount))
		args = append(args, filtros.Sexo)
		paramCount++
	}

	// Construcción final de la consulta
	query := fmt.Sprintf(baseQuery, selectFields)
	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	// Añadir agrupación según el tipo de gráfico
	switch tipo {
	case "barras", "lineas":
		query += " GROUP BY nacionalidad ORDER BY value DESC"
	case "pie":
		query += " GROUP BY sexo"
	}

	return query, args
}

// esta funciona bien con un query base
// func buildQuery(tipo string, filtros Filtros) (string, []interface{}) {
// 	var args []interface{}
// 	var conditions []string
// 	paramCount := 1

// 	// Base de la consulta según el tipo de gráfico
// 	baseQuery := `
// 		SELECT
// 			%s as name,
// 			COUNT(*) as value,
// 			AVG(%s) as extra_value
// 		FROM comparativas
// 		WHERE 1=1 `
// 	conditions = append(conditions, "nacionalidad !=''")
// 	// Definir las columnas específicas según el tipo de gráfico
// 	// aca le paso las variables que deseo calcular para cada grafico
// 	var nameColumn, extraColumn string
// 	switch tipo {
// 	case "barras":
// 		nameColumn = "nacionalidad"
// 		extraColumn = "idestablecimiento"
// 	case "lineas":
// 		nameColumn = "tipodependencia"
// 		extraColumn = "idencuesta"
// 	case "pie":
// 		nameColumn = "sexo"
// 		extraColumn = "idencuesta"
// 	default:
// 		// Manejo de tipos no válidos
// 		return "", nil
// 	}

// 	// Añadir filtros dinámicamente
// 	if filtros.Nacionalidad != "" {
// 		conditions = append(conditions, "nacionalidad = $"+strconv.Itoa(paramCount))

// 		args = append(args, filtros.Nacionalidad)
// 		paramCount++
// 	}
// 	if filtros.Educacion != "" {
// 		conditions = append(conditions, "tipodependencia = $"+strconv.Itoa(paramCount))
// 		args = append(args, filtros.Educacion)
// 		paramCount++
// 	}
// 	if filtros.Sexo != "" {
// 		conditions = append(conditions, "sexo = $"+strconv.Itoa(paramCount))
// 		args = append(args, filtros.Sexo)
// 		paramCount++
// 	}

// 	// Combinar la consulta base con las condiciones
// 	query := fmt.Sprintf(baseQuery, nameColumn, extraColumn)
// 	if len(conditions) > 0 {
// 		query += " AND " + strings.Join(conditions, " AND ")
// 	}

// 	// Añadir agrupación y ordenación según el tipo de gráfico
// 	switch tipo {
// 	case "barras", "lineas":
// 		query += " GROUP BY " + nameColumn + " ORDER BY value DESC"
// 	case "pie":
// 		query += " GROUP BY " + nameColumn
// 	}

// 	log.Println(query)
// 	log.Println(args)
// 	return query, args
// }

// Obtener datos del gráfico
func obtenerDatosGrafico(ctx context.Context, tipo string, filtros Filtros) ([]GraficoDatos, error) {
	query, args := buildQuery(tipo, filtros)

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var datos []GraficoDatos
	for rows.Next() {
		var dato GraficoDatos
		var extra float64 // Para el valor adicional (promedio_salario, promedio_edad o antiguedad_promedio)

		if err := rows.Scan(&nacionalidad, &dato.Value, &extra); err != nil {
			return nil, err
		}

		// Añadir información adicional al nombre según el tipo de gráfico
		switch tipo {
		case "barras":
			dato.Name = nacionalidad.String + " (Salario Prom: $" + formatFloat(extra) + ")"
		case "lineas":
			dato.Name = nacionalidad.String + " (Edad Prom: " + formatFloat(extra) + ")"
		case "pie":
			dato.Name = dato.Name + " (Años Prom: " + formatFloat(extra) + ")"
		}

		datos = append(datos, dato)
	}

	return datos, nil
}

// Formatear números flotantes
func formatFloat(num float64) string {
	return fmt.Sprintf("%.2f", num)
}

// Función para generar un gráfico específico
func generarGrafico(ctx context.Context, tipo string, filtros Filtros, resultChan chan<- Grafico, wg *sync.WaitGroup) {
	defer wg.Done()

	// Generar clave para el caché
	cacheKey := fmt.Sprintf("%s-%s-%s-%s", tipo, filtros.Nacionalidad, filtros.Educacion, filtros.Sexo)

	// Intentar obtener del caché
	if cached, found := c.Get(cacheKey); found {
		if datos, ok := cached.([]GraficoDatos); ok {
			resultChan <- Grafico{ID: tipo, Datos: datos}
			return
		}
	}

	// Si no está en caché, obtener de la base de datos
	datos, err := obtenerDatosGrafico(ctx, tipo, filtros)
	if err != nil {
		log.Printf("Error al obtener datos para gráfico %s: %v", tipo, err)
		resultChan <- Grafico{ID: tipo, Datos: []GraficoDatos{}}
		return
	}

	// Guardar en caché por 5 minutos
	// c.Set(cacheKey, datos, 5*time.Minute)
	// cache.NoExpiration -- si expiracion
	c.Set(cacheKey, datos, 5*time.Minute)

	resultChan <- Grafico{ID: tipo, Datos: datos}
}

// Función para generar todos los gráficos concurrentemente
func generarGraficos(ctx context.Context, filtros Filtros) []Grafico {
	tiposGraficos := []string{"barras", "lineas", "pie"}
	resultChan := make(chan Grafico, len(tiposGraficos))
	var wg sync.WaitGroup

	for _, tipo := range tiposGraficos {
		wg.Add(1)
		go generarGrafico(ctx, tipo, filtros, resultChan, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var graficos []Grafico
	for grafico := range resultChan {
		graficos = append(graficos, grafico)
	}

	return graficos
}

// Manejador HTTP para los gráficos
func handleGraficos(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var filtros Filtros
	if err := json.NewDecoder(r.Body).Decode(&filtros); err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	graficos := generarGraficos(ctx, filtros)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(graficos)
}

func closeDB() {
	if err := db.Close(); err != nil {
		log.Printf("Error al cerrar la base de datos: %v", err)
	}
}

func main() {
	// Inicializar caché con 5 minutos de expiración por defecto y limpieza cada 10 minutos
	c = cache.New(5*time.Minute, 10*time.Minute)

	// Inicializar base de datos
	if err := initDB(); err != nil {
		log.Fatal("Error inicializando base de datos:", err)
	}
	defer db.Close()

	// Configurar rutas
	http.HandleFunc("/graficos", handleGraficos)
	http.Handle("/", http.FileServer(http.Dir(".")))

	// Iniciar servidor
	log.Println("Servidor iniciado en :8090")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
