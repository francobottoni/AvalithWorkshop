package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		go imprimir(1, i)
		go imprimir(2, i)
		go imprimir(3, i)
	}

	time.Sleep(3 * time.Second)
}

func imprimir(numberGoroutine int, i int) {
	fmt.Println("Goroutine: ", numberGoroutine, "numero: ", i)
}
