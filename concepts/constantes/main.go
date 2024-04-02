package main

import "fmt"

const (
	os     int = 9
	domain int = 10
)

const (
	//operador iota contabiliza el grupo de constantes empezando en 0
	jan = iota + 1
	feb
	march
	april
	may
	june
)

func main() {

	// A diferencia de las variables el operador de asignacion corta := no
	// funciona con constantes
	// Ya que este asignara siempre a variables y no a constantes

	fmt.Println(april)
}

/*
	Nota: Las constantes no necesitan ser utilizadas en su totalidad
	a diferencia de las variables
 	El codigo compilara si tenemos constantes sin utilizar a diferencia
	de var sin utilizar
*/
