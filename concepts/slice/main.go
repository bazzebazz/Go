package main

import "fmt"

func main() {
	things := [6]string{"💅", "🚗", "🚓", "🚙", "🥎", "🏴‍☠️"}
	cars := things[:4]
	all := things[:]
	cars[0] = "🏎️"
	yellow := things[4]
	fmt.Println("Things: ", things)
	fmt.Println("Cars: ", cars)
	fmt.Println("Todo: ", all)
	fmt.Println("Yellows: ", yellow)
	// Cars[0] hace referencia al tamaño del slice. No al total del arreglo de things
	fmt.Println("Cars[0]: ", cars[0])
}
