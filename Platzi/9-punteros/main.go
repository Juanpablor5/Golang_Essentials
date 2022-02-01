package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {

	a := 50
	b := &a

	fmt.Println("a", a)
	fmt.Println("Dirección en memoria de a:", b)
	fmt.Println("Valor de la dirección en memoria de a:", *b)

	// Aquí le estamos cambiando también el valor a "a", porque estamos cambiando el valor que guarda esa dirección de memoria, que es la de "a"
	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
}
