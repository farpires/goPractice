package main

import "fmt"

func main() {
	// fmt.Println(uno(5))
	// numero, estado := dos(2)
	// fmt.Println(numero)
	// fmt.Println(estado)
	fmt.Println(Calculo(5, 46))
	fmt.Println(Calculo(2, 23, 45, 68))
	fmt.Println(Calculo(5))
	fmt.Println(Calculo(5, 16, 23, 46, 94, 98))
}

func uno(numero int) (z int) {
	z = numero * 2
	return
}

func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

func Calculo(numero ...int) int {
	total := 0
	for item, num := range numero { //rango de todod los elemento que tenemos en mi lista de nuemero . ram intursio de g que devuelve dos valores
		total = total + num
		fmt.Printf("\n Item %d \n", item)
	}
	return total
}
