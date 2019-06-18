package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/edwinbaeza/gocurso/estructuras"
	"github.com/edwinbaeza/gocurso/goruntimes"
	"github.com/edwinbaeza/gocurso/interfaces"
	"github.com/edwinbaeza/gocurso/maps"
	"github.com/edwinbaeza/gocurso/name"
	"github.com/edwinbaeza/gocurso/numbers"
	"github.com/edwinbaeza/gocurso/structs"

	//"github.com/edwinbaeza/gocurso/structs"
	"github.com/edwinbaeza/gocurso/Pointers"
)

const helloworld string = "Hola %s Bienvenido a go \n"

func main() {
	fmt.Println("Iniciando")
	var person *estructuras.Person
	if person == nil {
		fmt.Println("Es igual a nil")
	}
	Pointers.TestPointers()
	interfaces.TestInterfaces()
	goruntimes.MainGoRuntimes()
	time.Sleep(100000 * time.Millisecond)

	// person.Name = "Edwin"
	// person.Edad = 22
	// //estructuras.Person{Name: "Edwin", Edad: 22}
	// person.Saludar("Hola,")
	// fmt.Println()
	// fmt.Println(person.Name)
}

func mainTemp() {
	otroStirng := "Edwin Baeza"
	var number = sum(100, 233)
	a, b, c := getVariables()
	number = number + a + b + c
	firstName := name.GetName()
	fmt.Printf(helloworld, firstName)
	fmt.Println(otroStirng)
	fmt.Println(structs.GetSlice()[0])
	forTest()
	whileTest()
	stringTest()
	numbers.SwitchTest()
	fmt.Println(maps.GetMap())
}

func getVariables() (int, int, int) {
	return 1, 2, 2
}
func sum(a int, b int) int {
	return a + b
}

func probandoCondicionales(cadena string) bool {

	if lenght := len(cadena); lenght >= 20 {
		fmt.Println("La cadena es muy larga")
	} else {
		fmt.Println("La de cadena es muy corta")
	}

	return false
}
func forTest() {

	for index := 0; index < 5; index++ {
		fmt.Printf("Contador %d \n", index)
	}
}

func whileTest() {
	var input int = 0
	fmt.Println("Dame el numero ")
	fmt.Scanf("%d", &input)

	for input != 0 {
		fmt.Println("El numero fue distinto a 0")
		fmt.Println("Dame el numero ")
		fmt.Scanf("%d", &input)
	}
}

func stringTest() {
	var saludo = "Hola Mundo, Con Go, Hola Edwin"
	saludo1 := strings.ToUpper(saludo)
	saludo2 := strings.ToLower(saludo)
	saludo3 := strings.Replace(saludo, "Hola", "Hello", -1)
	cadena4 := strings.Split(saludo, ",")
	fmt.Println(saludo)
	fmt.Println(saludo1)
	fmt.Println(saludo2)
	fmt.Println(saludo3)
	fmt.Println(cadena4)
}
