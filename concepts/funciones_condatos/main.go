package main

import "fmt"

func main() {
	// Utilizamos la funcion para reutilizar con otras funciones
	// nums := []int{19, 22, 3, 44, 52, 1002}
	// result := filter(nums, lessThanTwenty)
	// fmt.Println(result)

	// result := sum(2)(3)
	// fmt.Println(result)

	// Asignamos el valor inicial en una variable para poder reutilizarla
	// Utilizamos la variable y enviamos otro valor para reutilizarla
	plus10 := sum(10)
	fmt.Println(plus10(2))
	fmt.Println(plus10(1))
	fmt.Println(plus10(4))
	fmt.Println(plus10(10))
}

// Funcion reutilizable que llama otras funciones con condiciones
// func greatherThanTwenty(num int) bool {
// 	return num > 20
// }
// func lessThanTwenty(num int) bool {
// 	return num < 20
// }
// func filter(nums []int, callback func(int) bool) []int {
// 	result := make([]int, 0, len(nums))
// 	for _, num := range nums {
// 		if callback(num) {
// 			result = append(result, num)
// 		}
// 	}
// 	return result
// }

//Funcion que retorna una funcion reutilizable de suma
func sum(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}
