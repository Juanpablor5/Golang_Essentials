package main

import (
	"fmt"
	"time"
)

func main() {
	numeros := make(chan int)
	cuadrados := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			numeros <- i
		}
		close(numeros)
	}()

	go func() {
		for {
			numero, open := <-numeros
			if !open {
				close(cuadrados)
				break
			}
			cuadrados <- numero * numero
		}
	}()

	for {
		data, open := <-cuadrados
		if !open {
			break
		}
		fmt.Println("1: ", data)
		time.Sleep(1 * time.Second)
	}

	// Ahora, una mejor forma de detectar si el chanel tiene datos, pero con el mismo ejemplo
	fmt.Println("========================")

	numeros2 := make(chan int)
	cuadrados2 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			numeros2 <- i
		}
		close(numeros2)
	}()

	go func() {
		for numero := range numeros2 {
			cuadrados2 <- numero * numero
		}
		close(cuadrados2)
	}()

	for data := range cuadrados2 {
		fmt.Println("2: ", data)
		time.Sleep(1 * time.Second)
	}

}
