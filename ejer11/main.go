package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leoArchivo()
	leoArchivo2()
	graboArchivo()
	graboArchivo2()
}

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		fmt.Println("Hubo un Error")
	} else {
		fmt.Println(string(archivo))
	}
}

func leoArchivo2() {
	archivo, err := os.Open("./file.txt")
	if err != nil {
		fmt.Println("Hubo un Error")
	} else {
		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			registro := scanner.Text()
			fmt.Printf("Linea > " + registro + "\n")
		}
	}
	archivo.Close()
}

func graboArchivo() {
	archivo, err := os.Create("./nuevoArchivo.txt")
	if err != nil {
		fmt.Println("Hubo un Error")
		return
	}
	fmt.Fprintln(archivo, "Esta es una line a nueva exitoooooo")
	archivo.Close()
}

func graboArchivo2() {
	fielName := "./nuevoArchivo.txt"
	if Append(fielName, "\n Esta es un segunda linea exito") == false {
		fmt.Println("Error en la segunda linea")
	}
}

func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Hubo un Error")
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un Error")
		return false
	}
	return true
}
