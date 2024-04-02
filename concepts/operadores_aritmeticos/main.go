package main

import "fmt"

func main() {
	// Operadores aritmeticos (), *, /, %, +, - esto en orden de ejecucion

	var a = (3 + 3) * 2
	fmt.Println(a)

	// Operador de asignacion: =, +=, -=, *=, /=, %=
	var b int = 5
	b += 2
	fmt.Println(b)

	// declaracion post-incremento y post-decremento: ++, --
	// (no son una expresion sino una declaracion)

	var c int = 8
	c++
	fmt.Println(c)
}
