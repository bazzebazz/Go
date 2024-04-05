package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type MyInt int
type MyIntV2 int

func main() {
	var num1 MyInt = 2
	var num2 MyInt = 3
	var num3 MyIntV2 = 2
	var num4 MyIntV2 = 3
	fmt.Println(sums[MyInt](num1, num2))
	fmt.Println(sums[MyIntV2](num3, num4))
	fmt.Println(sums[float32](2.2, 4.4, 52.2))
}

// type Number interface {
// 	~int | ~float32 | ~uint
// }

func sums[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}
