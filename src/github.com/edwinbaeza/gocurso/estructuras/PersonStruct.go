package estructuras

import (
	"fmt"
)

type Person struct {
	Name string
	Edad int16
}
// Saluda a una persona
func (p Person) Saludar(mensaje string) {

	fmt.Printf("%s nombre: %s de edad %d", mensaje, p.Name, p.Edad)
}
