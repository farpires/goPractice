package main

import "fmt"

type SerVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
	estaVivo() bool
}

type vegetal interface {
	ClasificacionVegetal() string
}

/*GENERO HUMANO*/
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}

type mujer struct {
	hombre
}

func (h *hombre) respirar() { h.respirando = true }
func (h *hombre) comer()    { h.comiendo = true }
func (h *hombre) pensar()   { h.pensando = true }
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	} else {
		return "Mujer"
	}
}
func (h *hombre) estaVivo() bool { return h.vivo }

func HumanoRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

/*-------------------------------------*/
/*GENERO ANIMAL*/

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (Pet *perro) respirar()         { Pet.respirando = true }
func (Pet *perro) comer()            { Pet.comiendo = true }
func (Pet *perro) EsCarnivoro() bool { return Pet.carnivoro }
func (Pet *perro) estaVivo() bool    { return Pet.vivo }

func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un Animal y estoy respirando")
}

func AnimalesCarnivoro(an animal) int {
	if an.EsCarnivoro() == true {
		return 1
	}
	return 0
}

/*Ser vivo*/

func estoyVivo(v SerVivo) bool {
	return v.estaVivo()
}

func main() {
	// Pedro := new(hombre)
	// Pedro.esHombre = true
	// HumanoRespirando(Pedro)

	// Maria := new(mujer)
	// HumanoRespirando(Maria)

	totalCarnivoro := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalesRespirar(Dogo)
	totalCarnivoro = +AnimalesCarnivoro(Dogo)
	fmt.Printf("total de carnivoro %d \n", totalCarnivoro)
	fmt.Printf(" Estoy vvivo = %t", estoyVivo(Dogo))

}
