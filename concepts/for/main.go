package main

import "fmt"

func main() {

	// for clasico
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// similar a un while
	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// For forever, se utiliza cuando necesitamos escuchar websockets o concurrencia
	// Podemos terminar el forever con una simple condicion que se tiene que cumpli
	// Y utlizar la palabra reservada break para salir del bucle
	// i := 1
	// for {
	// 	if i == 6 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	// food := []string{"🍕", "🍔", "🍍", "🌮", "🌭"}
	// for i, v := range []string{"🍕", "🍔", "🍍", "🌮", "🌭"} {
	// 	fmt.Println("Indice:", i, "Valor: ", v)
	// }

	// numbers := []uint8{2, 4, 6, 8}
	// for i := range numbers {
	// 	numbers[i] *= 2
	// }
	// fmt.Println(numbers)

	// food := map[string]string{
	// 	"pizza":      "🍕",
	// 	"hamburguer": "🍔",
	// 	"apple":      "🍎",
	// 	"hotdog":     "🌭",
	// }
	// for key, value := range food {
	// 	fmt.Println("Key: ", key, "Valor: ", value)
	// }

	for i, v := range "Nalgona" {
		// De esta forma muestra el valor unicode del formato Rune
		fmt.Println("Indice: ", i, "Valor: ", v)
		// De esta forma formatea el valor unicode a string con su valor
		fmt.Println("Indice: ", i, "Valor: ", string(v))
	}
}
