package main

import (
	"fmt"
	"time"
)

func main() {
	const n = 45
	go printAnimation()
	result := fibonacciSucession(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, result)
}

func printAnimation() {
	for {
		for _, c := range `\|/-` {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("\r%c", c)
		}
	}
}

func fibonacciSucession(n int) int {
	if n < 2 {
		return n
	}
	return fibonacciSucession(n-1) + fibonacciSucession(n-2)
}
