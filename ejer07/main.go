package main

import "fmt"

//array y vector no es mas que nada , que una coleccion de datos en memoria

//FORMA 1
// var tabla [10]int
// func main() {
// 	tabla[0] = 1
// 	tabla[5] = 15
// 	fmt.Println(tabla)
// }

//FORMA2
// func main() {
// 	tabla := [10]int{5, 6, 8, 9, 1, 23, 65, 3, 7, 45}
// 	fmt.Println(tabla)
// }

//FORMA3
// func main() {
// 	tabla := [10]int{5, 6, 8, 9, 1, 23, 65, 3, 7, 45}
// 	for i := 0; i < len(tabla); i++ {
// 		fmt.Println(tabla[i])
// 	}
// }

// FORMA4 - MATRIz
// var matriz [5][7]int
// func main() {
// 	matriz[3][5] = 1
// 	fmt.Println(matriz)
// }

//FORMA5  VECTORES DINAMICOS

func main() {
	// matriz := []int{2, 5, 4}
	// fmt.Println(matriz)
	variante2()
	variante3()
	variante4()
}
func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[3:] // te crea un vector apartir de la posiion 3 o porcion := elementos[3:5]

	fmt.Println(porcion)
}

// make tienen dos parametro (tipo, tamaño, cantidad)
//me reserva el pacio del vector pero te crea un vector de 5(tamaño)
//  lo mas costoso que un puede hacer es concatenar dos string en memoria
// por la cantida de operacion que tienen que hacer
// lo que se penso es que se puede dar una cantidad maxima , seria como el varchar
func variante3() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Cantidad %d", len(elementos), cap(elementos))
}

//appen tiene su costo de proceso
// para go es milesima de segundo
//
func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))

}
