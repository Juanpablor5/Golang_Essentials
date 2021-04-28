package main

import "fmt"

func main() {
	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 25,
		"Bayern":      75}

	campeonato["River Plate"] = 25 // Adiciona un item

	delete(campeonato, "Real Madrid") // Elimino un item

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de: %d\n", equipo, puntaje)
	}

	fmt.Println(campeonato)

	puntaje, existe := campeonato["Millonarios"]
	fmt.Printf("Puntaje: %d, existe equipo: %t", puntaje, existe)
}

func mapMake() {
	/*
	* Creo un map, asignandole los index de tipo string [string] y valores de tipo string también
	 */
	paises := make(map[string]string)

	paises["Mexico"] = "D.F."
	fmt.Println(paises["Mexico"])
}
