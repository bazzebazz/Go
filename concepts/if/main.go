package main

import "fmt"

func main() {
	// character := "🙆"
	if character := "🦸"; character == "🦸" {
		fmt.Println("Superhero")
	} else if character == "🦹" {
		fmt.Println("Supervillian")
	} else {
		fmt.Println("Civil")
	}
	// si utilizamos la declaracion de la variable dentro del scope If
	// esta varaible deja de estar disponible fuera del scope
	// fmt.Println(character)
}
