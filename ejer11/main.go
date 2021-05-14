package main

type humano interface {
	respirar()
	pensar()
	comer()
	Sexo() string
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() string
}

type vegetal interface {
	ClasificacionVegetal() string
}

// Genero Humano
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
}

type mujer struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
}

func main() {

}
