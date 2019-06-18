package estructuras

import (
	"errors"
	"fmt"
)

//basada en un tipo ya existente
//type MyStruct int

type PlatziCourse struct {
	Name   string
	Slug   string
	Skills []string
}

type OtherStruc struct {
	//tipo herencia
	PlatziCourse
}

func main() {
	platziCourse := PlatziCourse{
		Name: "Go", Slug: "go", Skills: []string{"1", "2"},
	}
	platziCourse2 := new(PlatziCourse)
	platziCourse2.Name = ""
	platziCourse2.Skills = []string{"backednd"}
	platziCourse2.Slug = ""
	fmt.Println(platziCourse)
	err := errores()
	fmt.Println(err)
	fmt.Println("The End")
}

func errores() error {
	return errors.New("Metodo Indefinidido")
}
