package main

import "fmt"

func main() {
	division(200, 10)
	division(200, 25)
	division(45, 0)
	division(75000, 4)
}

func division(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ya paso el panico, bruh")
		}
	}()

	validate0(divisor)
	fmt.Println(dividendo / divisor)
}

func validate0(divisor int) {
	if divisor == 0 {
		panic("😱 No puedes didivir por 0, que te pasa")
	}
}
