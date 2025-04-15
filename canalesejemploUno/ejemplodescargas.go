package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Archivo struct {
	URL       string
	Contenido string
}

func descargarArchivo(id int, urls <-chan string, resultados chan<- Archivo, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range urls {
		fmt.Printf("[Worker %d] Descargando desde: %s\n", id, url)
		// Simula una descarga con una espera aleatoria
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		resultados <- Archivo{URL: url, Contenido: "Contenido simulado"}
		fmt.Printf("[Worker %d] Descarga completada: %s\n", id, url)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	urls := []string{
		"https://sitio.com/archivo1",
		"https://sitio.com/archivo2",
		"https://sitio.com/archivo3",
		"https://sitio.com/archivo4",
		"https://sitio.com/archivo5",
	}

	trabajos := make(chan string, len(urls))
	resultados := make(chan Archivo, len(urls))
	var wg sync.WaitGroup

	numTrabajadores := 3
	for i := 1; i <= numTrabajadores; i++ {
		wg.Add(1)
		go descargarArchivo(i, trabajos, resultados, &wg)
	}

	for _, url := range urls {
		trabajos <- url
	}
	close(trabajos)

	wg.Wait()
	close(resultados)

	fmt.Println("\n✅ Archivos descargados:")
	for archivo := range resultados {
		fmt.Printf("URL: %s, Contenido: %s\n", archivo.URL, archivo.Contenido)
	}
}
