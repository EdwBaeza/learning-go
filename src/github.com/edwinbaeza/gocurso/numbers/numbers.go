package numbers

import "fmt"

//calcula si un numero es para o impar
func ParOrImpar(dato int) bool {

	if dato%2 == 0 {
		return true
	}
	return false
}

// prueba del switch
func SwitchTest() {

	var numero int
	fmt.Scanf("%d", &numero)
	switch numero {
	case 1:
		fmt.Println("El numero es uno")
		break
	default:
		fmt.Println("El numero no es uno")

	}

	switch {
	case numero%2 == 0:
		fmt.Println("El numero es par")
		break
	default:
		fmt.Println("El numero es impar")
	}
}
