package main

import "fmt"

func main() {
	// Tamaño de elementos
	// len(): numero de elementos en el slice
	// cap(): numero de elementos del array de origen, apartir del indice donde se creo
	// el slice
	// animals := [5]string{"🦍", "🦝", "🐴", "🐶", "🐦"}
	// pets := animals[1:3]

	// pets = append(pets, "🐈", "🐘")
	// Si sobrepasamos el valor del tamaño del array con append, este nos creara un nuevo array
	// Con el doble de capacidad es decir, en este caso la capacidad es de 4, si sobrepasamos
	// La capacidad seria de 8

	// fmt.Println("Animals: ", animals)
	// fmt.Println("Pets: ", pets)

	// fmt.Println("Tamaño de pets: ", len(pets))
	// fmt.Println("Capacidad de pets: ", cap(pets))

	// pets := []string{"🐈", "🐘"}

	// pets := make([]string, 0, 3)
	// pets = append(pets, "🐈", "🐘", "🐴")

	var pets []string
	fmt.Println("Pets: ", pets)

	fmt.Println("Tamaño de pets: ", len(pets))
	fmt.Println("Capacidad de pets: ", cap(pets))
	fmt.Println("Valor cero: ", pets == nil)
}
