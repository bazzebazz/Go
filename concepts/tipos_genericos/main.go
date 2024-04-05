package main

import "fmt"

type Product[T uint | string] struct {
	ID          T
	Description string
	Price       float64
}

func main() {
	product1 := Product[uint]{ID: 1, Description: "Zapato", Price: 7.99}
	product2 := Product[string]{ID: "asdsada212d12-12312dassa", Description: "Camisa", Price: 71.99}
	fmt.Println(product1)
	fmt.Println(product2)
}
