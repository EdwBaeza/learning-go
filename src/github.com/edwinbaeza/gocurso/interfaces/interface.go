package interfaces

import (
	"fmt"
)

type Interfaz interface {
	Saludar()
}

type Persona struct {
	Name string
	Edad int
}

func (p Persona) Saludar() {
	fmt.Printf("Hola Persona %s \n", p.Name)
}

type Alumno struct {
	Persona
	Matricula string
}

func (p Alumno) Saludar() {
	fmt.Printf("Hola Alumno %s \n", p.Name)
}

func SaludarGenerico(inter Interfaz) {
	inter.Saludar()
}

func TestInterfaces() {

	persona := Persona{Name: "Edwin", Edad: 22}
	alumno := new(Alumno)
	alumno.Name = "David"
	alumno.Edad = 23
	alumno.Matricula = "15080299"

	persona.Saludar()
	alumno.Saludar()

}
