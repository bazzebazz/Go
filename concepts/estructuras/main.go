package main

import "fmt"

type Person struct {
	Name        string
	Age         uint8
	HasChildren bool
}

func main() {
	alex := Person{
		Name:        "Alexander",
		Age:         26,
		HasChildren: true,
	}

	javier := Person{"Javier", 26, false}

	camila := Person{Age: 24}

	andrea := &alex
	andrea.Age = 27

	fmt.Printf("%+v \n", alex.Name)
	fmt.Printf("%+v \n", javier.HasChildren)
	fmt.Printf("%+v \n", camila)

	fmt.Println("*-*********-********-*")

	fmt.Printf("%+v \n", alex)
	fmt.Printf("Andrea %+v", andrea)
}
