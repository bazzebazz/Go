package main

import (
	"fmt"
	"strings"
)

func main() {
	// greetings("Javier", "Urrutia")
	name := "Alexander"
	// Aca mandamos el puntero para poder acceder al valor en la funcion
	toUpperCase(&name)

	fmt.Println(name)
}

// funcion sencilla tipica con parametros
// func greetings(firstName string, lastName string) {
// 	fmt.Println("Hello ", firstName, lastName)
// }

// Para que este codigo funcione fue necesario utilizar punteros y ademas utilzar el simbolo
// Anspersan (No se como se escribe) en el parametro que mandamos a la funcion
func toUpperCase(text *string) {
	*text = strings.ToUpper(*text)
}
