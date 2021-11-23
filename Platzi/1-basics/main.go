package main

import "fmt"

func main() {
	// Constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	// Variables
	var altura int = 14
	base := 12
	var area int

	fmt.Println(altura, base, area)

	// Area cuadrado
	areaCuadrado := base * altura
	fmt.Println("area cuadrado", areaCuadrado)

	// FMT
	s1 := "Hello"
	s2 := "World"
	fmt.Println(s1, s2)
	fmt.Printf("%s %s\n", s1, s2)

	// Sprintf: guarda como string
	s3 := fmt.Sprintf("%s %s\n", s1, s2)
	fmt.Println(s3)

	// Conocer el tipo de dato
	fmt.Printf("type s1: %T\n", s1)

	// Funciones
	result1, _ := firstFunction(altura, base, s1)
	fmt.Println(result1)
}

func firstFunction(a int, b int, c string) (string, int) {
	fmt.Println(a, b, c)
	return fmt.Sprintf("%s, el area es %d", c, a*b), a
}
