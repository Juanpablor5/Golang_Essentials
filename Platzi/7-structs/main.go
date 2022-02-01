package main

import "fmt"

type car struct {
	brand string
	year  int
}

// Struct tiene un método llamado String, que podemos sobrescribir para personalizar la salida a consola de los datos del struct.
func (myCar car) String() string {
	return fmt.Sprintf("El carro es de marca %s y modelo %d", myCar.brand, myCar.year)
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Otra manera
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)

}
