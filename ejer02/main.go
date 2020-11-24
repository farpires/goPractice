package main

import (
	"fmt"
)

var numero int
var texto string
var status bool = true

func main() {

	numero2, numero3, numero4, texto, status := 2, 5, 67, "Este es mi texto", false

	var moneda int64 = 0
	numero2 = int(moneda)
	// texto = fmt.Sprintf("%d", moneda) //manera de combertir
	// texto = strconv.Itoa(int(moneda)) //2da manera de combertir un string en texto
	fmt.Print(numero2)
	fmt.Print(numero3)
	fmt.Print(numero4)
	fmt.Print(texto)
	fmt.Print(status)
	mostrarStatus()
}

func mostrarStatus() {
	fmt.Println(status)
}
