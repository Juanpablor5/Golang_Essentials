package main

import "fmt"

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

// Función con un argumento receptor
func (c cuadrado) area() float64 {
	return c.base * c.base
}

// Función con un argumento receptor
func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2D) {
	fmt.Println("Area:", f.area())
}

func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	calcular(myCuadrado)
	calcular(myRectangulo)

	// Lista interfaces, para tener varios tipos de datos
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)
}
