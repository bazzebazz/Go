package main

import (
	"fmt"
	"os"
)

func main() {
	// defer nos servira para cerrar conexiones o bloquearlas cuando se terminen de ejecutar
	// Lo que nos ayudara mucho en el rendimiento de la aplicacion
	// num := 4
	// defer fmt.Println("Numero: ", num)
	// defer fmt.Println(2)
	// defer fmt.Println(1)
	// num = 10
	// fmt.Println("Numero: ", num)

	file, err := os.Create("text.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hello gophersssssssss"))
	if err != nil {
		fmt.Println(err)
		return
	}
}
