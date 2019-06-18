package Listas

import (
	"fmt"
)
type Nodo struct {
	Data string
	Liga *Nodo
}
type ListaLigada struct {
	Primero *Nodo
	Ultimo  *Nodo
}
func (list *ListaLigada) Agregar(dato string) {
	var nodo Nodo
	nodo.Data = dato
	//	var apuntador *Nodo
	if list.Primero == nil {
		nodo.Liga = nil
		list.Primero = &nodo
	} else {
		list.Ultimo.Liga = &nodo
	}
	list.Ultimo = &nodo
}

func (list ListaLigada) Recorrer() {

	temporal := list.Primero
	for temporal != nil {
		fmt.Println(temporal.Data)
		temporal = temporal.Liga
	}
}
