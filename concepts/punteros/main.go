package main

import "fmt"

func main() {
	// puntero: variables que almacenan la direccion en memoria de un valor
	var color string = "🟥"
	// *string o algun tipo de dato hace un puntero a una variable de ese tipo
	var pointerColor *string
	pointerColor = &color
	*pointerColor = "🟦"

	fmt.Printf("Tipo: %T Valor: %s, Direccion: %v\n", color, color, &color)
	fmt.Printf("Tipo: %T Valor: %v, Desreferencia: %s\n", pointerColor, pointerColor, *pointerColor)

}
