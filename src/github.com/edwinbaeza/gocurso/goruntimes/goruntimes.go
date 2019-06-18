package goruntimes

import (
	"fmt"
	"time"
)

func Imprimir(num int) {

	for i := 0; i < num; i++ {

		fmt.Printf("Numero %d de un total de %d \n", i, num)
		numRandom := 1000 * time.Millisecond
		time.Sleep(numRandom)
	}
}

func MainGoRuntimes() {
	go Imprimir(100)
	go Imprimir(150)
}
