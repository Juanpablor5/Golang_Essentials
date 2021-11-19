package main

import (
	"fmt"
)

var numero int // por defecto se inicializa en 0, "", false; según sea.
var texto string
var status bool

func main() {
	numero5 := 4 //automáticamente me da el tipo con base en la asignación
	var numero, numero2, numero3 int
	// var status bool
	numero, numero2, numero3, status = 1, 2, 3, true
	obtenerStatus()

	var moneda int64 = 0

	numero2 = int(moneda)
	texto = fmt.Sprintf("%d", moneda)

	fmt.Println(numero)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero5)
	fmt.Println(numero5)
	fmt.Println(texto)
	fmt.Println(status)

}

func obtenerStatus() {
	fmt.Println(status)
}
