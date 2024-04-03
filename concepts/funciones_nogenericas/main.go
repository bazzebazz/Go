package main

import "fmt"

func main() {
	PrintList("Testing", "Gopher", "Go desde 0")
	PrintList(1, 2, 3)
}

func PrintList(list ...any) {
	for _, item := range list {
		fmt.Println(item)
	}
}
