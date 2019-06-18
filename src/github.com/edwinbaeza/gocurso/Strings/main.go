package main

import (
	"fmt"
	"strings"
)

func main() {
	capitalize("Hola Mundo")
}

func capitalize(str string) {

	withoutLastLetter := str[:len(str)-1]
	fmt.Println(withoutLastLetter)
	lastLetter := string(str[len(str)-1])
	withoutLastLetter += strings.ToUpper(lastLetter)
	fmt.Println(withoutLastLetter)
}
