package main

import (
	"fmt"
	"github.com/edwinbaeza/Algoritmos/list"
)

func main() {
	fmt.Println("Inicio de Insercion")
	lista := new(list.LinkedList)

	lista.Add("Edwin")
	lista.Add("David")
	lista.Add("Virgilio")
	lista.PrintList()
}
