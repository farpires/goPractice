package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var resultado int
var leyenda int

func main() {
	fmt.Println("Ingrese numero 1:")
	fmt.Scanln(&numero1)
	fmt.Println("Ingrese numero 2:")
	fmt.Scanln(&numero2)
	fmt.Printf("%d", numero1+numero2)
	fmt.Println("Descripcion: ") //introduccion de cabecera
	//bufio es para lectura de archivo lectura por teclado

	scanner := bufio.NewScanner(os.Stdin)
	//stdin = se imprime en la pantalla
	// esto es una manera de controlar si es que no viene con error
	if scanner.Scan() {
		leyenda = scanner.Text()
	}
	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)
}
