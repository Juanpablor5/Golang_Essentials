package main

import "fmt"

var array [10]int
var matrix [5][7]int

func main() {
	variante2()
	variante3()
	variante4()
}

func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[:3] // Me retorna desde el elemento 0 hasta el 2
	fmt.Println(porcion)
}

func variante3() {
	elementos := make([]int, 5, 20) // Crea un slice de tamaño 5 con capacidad de 20, pero me reserva el espacio sin usarlo
	fmt.Printf("Largo %d, Capacidad %d\n", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)
	// nums := []int{} // También se puede crear inferido

	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Println(nums)
	fmt.Printf("Largo %d, Capacidad %d\n", len(nums), cap(nums))
}

func slice() {
	matriz := []int{2, 5, 4} // Si no le asigno longitud, me infiere que es un Slice, es decir, un arreglo de tamaño variable
	fmt.Println(matriz)
}

func matrices() {
	matrix[2][5] = 5

	fmt.Println(matrix)
}

func arreglos() {
	tabla := [10]int{5, 16, 2, 5, 1, 4}

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}

	fmt.Println(tabla)
}
