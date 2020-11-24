package main

import "fmt"

func main() {
	// paises := make(map[string]string, 5)
	// fmt.Println(paises)

	// paises["Mexico"] = "D.F."
	// paises["Argentina"] = "Catamarca"

	campeonato := map[string]int{
		"Barcelona":   39,
		"Real Madril": 38,
		"Cali":        35,
		"Boca Junior": 30}
	campeonato["River Plate"] = 25
	campeonato["Cali"] = 55
	delete(campeonato, "Real Madril")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El Equipo %s, tiene un puntaje de : %d \n", equipo, puntaje)
	}
	// en go no exite los null  ,,, cuando un valor inicializa y no existe devuelve un valor de cero
	puntaje, existe := campeonato["Cali"]
	fmt.Printf("el puntaje captura es %d , y el equipo existe %t \n", puntaje, existe)
}
