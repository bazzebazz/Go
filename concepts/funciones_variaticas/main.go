package main

import "fmt"

func main() {
	// fmt.Println(sum(2, 3))
	// fmt.Println(sum(2, 3, 12))
	// fmt.Println(sum(2, 3, 12, 1))
	// fmt.Println(sum(2, 3, 12, 1, 24))

	// Funcion anonima asignada a  una variable
	// greeting := func() {
	// 	fmt.Println("✋ Hello gopher")
	// }
	// greeting()

	// Funcion anonima autoinvocada con parametros
	func(name string) {
		fmt.Println("🤭 Hello gopher", name)
	}("Alex")
}

// Funcion variatica
/*func sum(nums ...int) (total int) {
	// var total int
	for _, num := range nums {
		total += num
	}
	return
}
*/
