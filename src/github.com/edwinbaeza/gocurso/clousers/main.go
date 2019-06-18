package main

import (
	"fmt"
)

func main() {
	funcion := func() {
		fmt.Println("Hola Mundo")
	}

	funcion()
}

func getFunc() func() {

	return func() {
		fmt.Println("Hola Mundo")
	}
}
