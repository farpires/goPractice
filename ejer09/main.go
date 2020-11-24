package main

//fecha = time
import (
	"fmt"
	"time"

	us "./user"
)

type pepe struct { //idem: extender , polimorfirmo, herencia
	us.Usuario //manera de hacer herencia
}

func main() {
	u := new(pepe)
	u.AltaUsuario(1, "Fabio Arpires", time.Now(), true)
	fmt.Println(u.Usuario)
}
