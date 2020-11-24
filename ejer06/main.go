package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 7 = %d \n", Calculo(5, 7))

	// Restamos
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Restamos 6 - 4 = %d \n", Calculo(6, 4))
	//dividimos
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("DIVIDIMOS 12/3 = %d \n", Calculo(12, 3))
	//multiplicamos
	Calculo = func(num1 int, num2 int) int {
		return num1 * num2
	}
	fmt.Printf("Multiplicamos 2*2 = %d \n", Calculo(2, 2))
	Operacion()
	/*CLOUSERES*/
	tablaDel := 2
	MiTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}

// funciones anonimas

func Operacion() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}

// devuelve otra funcion de tipo entero
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia += 1
		return numero + secuencia
	}
}
