package main

import (
	"fmt"

	"github.com/edwinbaeza/Algoritmos/Listas"
)

func main() {
	fmt.Println("Inicio de Insercion")
	lista := new(Listas.ListaLigada)

	// if lista.Primero == nil {
	// 	fmt.Println("Si es nil primero")
	// }
	lista.Agregar("Edwin")
	//fmt.Println(lista.Primero.Data)
	lista.Agregar("David")
	lista.Agregar("Virgilio")
	lista.Recorrer()
}
