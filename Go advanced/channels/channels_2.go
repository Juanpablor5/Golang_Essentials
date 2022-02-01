// undirectional channels
package main

import (
	"fmt"
)

func main() {
	nc := make(chan int)
	sc := make(chan int)

	go numberGenerator(nc)
	go squaring(nc, sc)

	printNumber(sc)
}

func numberGenerator(out chan<- int) {
	for i := 0; i < 5; i++ {
		out <- i
	}
	close(out)
}

func squaring(in <-chan int, out chan<- int) {
	for n := range in {
		out <- n * n
	}
	close(out)
}

func printNumber(in <-chan int) {
	for n := range in {
		fmt.Println(n)
	}
}
