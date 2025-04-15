package main

import (
	"fmt"
	"sync"
)

func trabajador(id int, trabajos <-chan int, resultados chan<- int) {
	for j := range trabajos {
		fmt.Printf("Trabajador %d procesando trabajo %d\n", id, j)
		resultados <- j * j
	}
}

func main() {
	const numTrabajadores = 3
	const numTrabajos = 9
	trabajos := make(chan int, numTrabajos)
	resultados := make(chan int, numTrabajos)
	var wg sync.WaitGroup

	for w := 1; w <= numTrabajadores; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			trabajador(id, trabajos, resultados)
		}(w)
	}

	for i := 1; i <= numTrabajos; i++ {
		trabajos <- i
	}
	close(trabajos)

	wg.Wait()
	close(resultados)

	for res := range resultados {
		fmt.Println(res)
	}
}
