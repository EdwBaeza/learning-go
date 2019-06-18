package name

import "fmt"

func GetName() string {
	var name = "test"
	fmt.Println("Escribe tu nombre: ")
	fmt.Scanf("%s", &name)

	for index := 0; index < len(name); index++ {
		fmt.Println(string(name[index]))
	}
	return name
}
