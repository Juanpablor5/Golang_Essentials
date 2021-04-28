package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	num1, num2 := 5, 7
	fmt.Printf("Sumo %d + %d = %d", num1, num2, Calculo(num1, num2))

	Calculo = func(num1, num2 int) int {
		return num1 - num2
	}

	fmt.Printf("\nResto %d - %d = %d\n", num1, num2, Calculo(num1, num2))

	/* CLOSURES */
	tablaDel := 2
	MiTabla := Tabla(tablaDel) // Sólo tomó la función que le regresó Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
