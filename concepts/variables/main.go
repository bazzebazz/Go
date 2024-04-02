package main

import "fmt"

func main() {

	// declaracion de variable fav
	// var(
	// 	apple string = "hola"
	// 	banana string = "hola"
	// 	orange string = "hola"
	// )

	// var apple, banana, orange string = "🍎", "🍌", "🍊"
	// Declaracion de variable corta
	apple, banana, orange := "🍎", "🍌", "🍊"

	// Se pueden modificar valores de variables antes creadas con el operador := SIEMPRE Y CUANDO haya una nueva variable para declaraar
	// De lo contrario se utilizaria solamente apple = "Valor a cambiar"
	apple, lemon := "manzanasss", "🍋"

	fmt.Println(apple, banana, orange, lemon)
}
