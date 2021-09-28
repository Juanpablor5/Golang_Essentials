package main

import "fmt"

func main() {
	fmt.Printf("primer log %d\n", uno(5))

	// Si no se quiere guardar uno de los n parámetros, se pone _
	numero, estado := dos(1)
	fmt.Println(numero, estado)

	// Múltiples parámetros
	fmt.Println(Calculo(4, 1))
	fmt.Println(Calculo(2, 4, 3, 4))
	fmt.Println(Calculo(400, 2, 3, 5, 4, 1))
	fmt.Println(Calculo(4, 1, 5, 1654, 654, 4))
}

func uno(numero int) (z int) {
	z = numero * 2
	return
}

func dos(numero int) (int, bool) { // Cuando se van a retornar más de dos parámetros, deben guardarse los tipos en parentesis
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func Calculo(numero ...int) (total int) { // Cuando no se sabe cuantos parámetros se van a recibir
	for i := range numero { // Range toma todos los elementos de un arreglo y devuelve dos valores, el index y el elemento
		total = total + numero[i]
	}
	return
}
