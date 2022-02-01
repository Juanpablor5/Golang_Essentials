package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go printQuantity("A")
	go printQuantity("B")

	wg.Wait()
	fmt.Println("finished")
}

func printQuantity(t string) {
	// El defer indica que esta linea se ejecutará hasta que haya terminado toda la función
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
		fmt.Printf("Quantity %d of %s\n", i, t)
	}
}
