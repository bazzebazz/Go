package main

import "fmt"

func main() {
	// operadores de comparacion <, >, ==, !=, <=, >=
	fmt.Println(4 > 7)
	fmt.Println((4 + 5) > 7)
	// fmt.Println(7 >= 7)
	// fmt.Println(7 == 7.0)
	fmt.Println(7 != 7.0)

	// operadores logicos
	var edad uint = 26
	fmt.Println("Es adulto: ", edad >= 18 && edad <= 60)
	fmt.Println("Es niño o anciano: ", edad < 18 || edad > 60)

	// operador logico binario
	fmt.Println("Operador logico binario: ", !(4 != 4))
}
