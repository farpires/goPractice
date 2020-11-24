package main

import "log"

func main() {
	// archivo := "prueba.txt"
	// f, err := os.Open(archivo)
	// defer f.Close()

	// if err != nil {
	// 	fmt.Println("error abriendo el archivo ")
	// 	os.Exit(1)
	// }

	ejemploPani()
}

func ejemploPani() {
	defer func() {
		//defer ejecuta una sola cosa un close
		//si ejecutas mas de una cosa utilizaremo f\uncions
		//recover trae el resultado de ultiima panic
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que genero panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("encontro el valor de 1")
	}
}
