package main

import "fmt"

func main() {
	i := 0
RUTINA:
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Println("Voy a RUTINA sumando 2 a i")
			goto RUTINA
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}
}

func continueFunction() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("\n d: %d", i)
		if i == 5 {
			fmt.Print(" Entra y continúa \n")
			continue // continúa a la siguiente iteración, finalizando aquí aquí
		}
		fmt.Print(" Continúa sin entrar\n")
	}
}
