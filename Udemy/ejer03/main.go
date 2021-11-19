package main

import "fmt"

var estado bool

func main() {
	estado = true
	if estado = false; estado { // Le puedo asignar un valor dentro del if
		fmt.Println("Hola")
	} else {
		fmt.Println("Adiós")
	}

	switch numero := 5; numero {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println(6)
	}
}
