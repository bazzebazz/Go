package main

import "fmt"

func main() {
	// bool, string, numeric
	var a bool = true
	var b string = "golang"
	// En los tipos de datos numericos hay muchos de tipos de dato
	// No solamente int
	var c uint8 = 8
	var d uint16 = 26
	var e rune = 'f'
	var f float32 = 125.99

	// hacer el casting nos permite realizar la operacion pero esto no significa
	// que cambiara el tipo de dato
	convertir := uint16(c) + d
	// si queremos compilar el codigo y mantener la variable para ellos hacemos esto
	// _ = uint16(c) + d operador blank
	// de esta manera go nos dejara compilar el codigo

	// documentacion de donde saco: https://pkg.go.dev/fmt@go1.22.1
	// %T	a Go-syntax representation of the type of the value
	// %v	the value in a default format
	fmt.Printf("Tipo: %T, Valor: %v \n", a, a)
	fmt.Printf("Tipo: %T, Valor: %v \n", b, b)
	fmt.Printf("Tipo: %T, Valor: %v \n", c, c)
	fmt.Printf("Tipo: %T, Valor: %v \n", d, d)
	fmt.Printf("Tipo: %T, Valor: %v \n", e, e)
	fmt.Printf("Tipo: %T, Valor: %v \n", f, f)
	fmt.Printf("Tipo: %T, Valor: %v", convertir, convertir)
}
