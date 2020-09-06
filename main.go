//al inicializar una variable sin ningun valor, boolean= false, int=0 testo=""
//scoot = es el alcance
//en go practicamente todo son funciones
package main

import (
	"fmt"
	"strconv"
)

//variable entero
var numero int

// numerotwo := 2
//variable string
var texto string
var textotwo string

//variable booleano
var status bool

func main() {
	var numero2, numero3, numero4 int
	numero2, numero3, numero4, texto, status := 2, 5, 67, "Este es mi texto", true
	//movimiento de entero de 64bit a  un entero-------------
	var moneda int64 = 64
	numero2 = int(moneda)
	//--------------------------------------------------------
	//convertir un numero a String ---------RECORDAR, tiene diferentes verbos par su correspondiente cambio-----------------
	var money int64 = 1500
	texto = fmt.Sprintf("%d", money)
	//--------------------------------------------------------
	//--usando strcov.itoa(int())

	textotwo = strconv.Itoa(int(money)) + "esta es la segunda conversion"

	//---------------
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(status)
	fmt.Println(textotwo)
	mostrarStatus()
}
func mostrarStatus() {
	fmt.Println(status)
}
