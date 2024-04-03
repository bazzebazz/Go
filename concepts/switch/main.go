package main

import "fmt"

func main() {
	character := "🦸"
	canSearch := false

	switch character {
	case "🦸", "🧞":
		fmt.Println("Superhero")
	// case "🧞":
	// 	fmt.Println("Superhero")
	case "🦹", "🧟":
		fmt.Println("Supervillian")
	// case "🧟":
	// 	fmt.Println("Supervillian")

	default:
		fmt.Println("Civil")
	}

	//Otra manera de evaluar un switch

	// De esta manera podemos evaluar diferentes casos a comparacion con el anterior

	switch {
	case !canSearch:
		fmt.Println("Busqueda no permitida")
	case character == "🦸" || character == "🧞":
		fmt.Println("Superhero")
	// case "🧞":
	// 	fmt.Println("Superhero")
	case character == "🦹" || character == "🧟":
		fmt.Println("Supervillian")
	// case "🧟":
	// 	fmt.Println("Supervillian")

	default:
		fmt.Println("Civil")
	}
}
