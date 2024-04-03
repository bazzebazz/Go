package main

import (
	"fmt"
	"strings"
)

func main() {
	// result := sum(2, 3)
	// fmt.Println(result)

	lower, upper := convert("Nalgona")

	fmt.Println(lower, upper)
}

// func sum(a int, b int) int {
// 	return a + b
// }

// esta forma de retorno es recomendable unicamente
// Cuando las funciones son cortas
func convert(text string) (lower string, upper string) {
	lower = strings.ToLower(text)
	upper = strings.ToUpper(text)

	return
}
