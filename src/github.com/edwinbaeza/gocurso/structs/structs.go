package structs

import "fmt"

//Obtiene un array
func GetArray() [2]int {
	var array [2]int
	//otherArray := [2]int{10, 20}
	primes := [6]int{2, 3, 5, 7, 11, 13}
	primes2 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(primes[0] + primes2[0])
	array[0] = 12
	array[1] = 200
	return array
}

// obtiene un slice
func GetSlice() []int {
	var slice []int
	slice = append(slice, 10, 20, 30, 40)
	slice = append(slice, 50, 60, 70, 80)
	return slice
}
