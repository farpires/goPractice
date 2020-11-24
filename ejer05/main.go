package main

import "fmt"

func main() {
	// 	i := 1
	// 	for i < 10 {
	// 		fmt.Println(i)
	// 		i++
	// 	}

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)

	// }
	// numero := 0
	// for {
	// 	fmt.Println("continuo")
	// 	fmt.Println("ingresa el numeor secreto")
	// 	fmt.Scanln(&numero)
	// 	if numero == 100 {
	// 		break
	// 	}
	// }
	// var i = 0
	// for i < 10 {
	// 	fmt.Printf("\n valor de i: %d", i)
	// 	if i == 5 {
	// 		fmt.Println("multiplicamos por 2 \n")
	// 		i = i * 2
	// 		continue
	// 	}
	// 	fmt.Printf(" paso  por aqui \n")
	// 	i++
	// }
	i := 0
RUTINA:
	for i < 10 {
		if i == 4 {
			i = i + 2
			fmt.Printf("Voy a rutina sumanod dos  %d\n", i)
			goto RUTINA
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}

}
