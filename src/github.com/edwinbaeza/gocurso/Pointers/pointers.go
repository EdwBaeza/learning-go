package Pointers

import (
	"fmt"
)

//probando apuntadores
func TestPointers() {

	var num1 = 100
	var apuntador *int
	apuntador = &num1

	fmt.Println(num1, *apuntador, &num1)

}
