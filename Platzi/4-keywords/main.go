package main

import "fmt"

func main() {
	// Defer: hace que esa línea se ejecute justo antes de terminar el código de la función
	defer fmt.Println("2do")
	fmt.Println("1ro")

	// Continue: funciona igual que en JS, continua la siguiente iteración si seguir en las lineas siguientes

	// Break: finaliza el loop
}
