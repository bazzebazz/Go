package main

import "fmt"

func main() {
	music := make(map[string]string)
	music["guitar"] = "🎸"
	music["violin"] = "🎻"

	fmt.Println(music)

	tech := map[string]string{
		"computer": "💻",
		"mouse":    "🖱️",
	}

	fmt.Println(tech)
	delete(tech, "computer")
	fmt.Println(tech)

	fmt.Println(music["guitar"])
	// si queremos acceder al arreglo de map de un valor que no existe, este
	// nos devolvera el valor 0 del string es decir un valor vacio ""
	content, ok := music["piano"]
	fmt.Println(content, ok)
}
