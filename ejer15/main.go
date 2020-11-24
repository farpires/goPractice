package main

import (
	"fmt"
	"time"
)

//el cananl transporta una variable
// Sub : mde retorna la duracion
func main() {
	//canal utilizado para calcular el calculo de ua diferencia
	//entre distintas hora . cauunto milisegundo dura laejecucion
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("llege hasta aqui")

	msg := <-canal1
	fmt.Println(msg)
}

func bucle(canal1 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 10000000000; i++ {

	}
	final := time.Now()
	canal1 <- final.Sub(inicio)
}
