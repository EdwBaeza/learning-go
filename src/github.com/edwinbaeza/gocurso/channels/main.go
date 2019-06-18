package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	channel := make(chan string)
	go generarValor(channel)
	fmt.Println("Inicio")
	fmt.Println(<-channel)
	fmt.Println("final")
	go imprimirValor(channel)
	fmt.Scanln()
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generarValor(canal chan string) {
	var str2 strings.Builder
	//nuevo := rand.Intn(1000)
	for index := 0; ; index++ {
		str2.WriteString("Ping ")
		str2.WriteString(strconv.Itoa(index))
		canal <- str2.String()
		str2.Reset()
	}

	//fmt.Println(str2.String())
}

func imprimirValor(canal chan string) {

	for {
		fmt.Println(<-canal)
		time.Sleep(time.Second * 1)
	}
}
