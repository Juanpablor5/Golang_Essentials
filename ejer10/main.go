package main

import (
	"fmt"
	"time"

	us "go/ejer10/user"
)

// La estructura de pepe, hereda de usuario
type pepe struct {
	us.Usuario
}

func main() {
	user := new(pepe)
	user.AltaUsuario(1, "Pablo Tilotta", time.Now(), true)
	fmt.Println(user.Usuario)
}
