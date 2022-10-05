package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Goroutines")
	go escrevendo("Olá Mundo") //goroutines
	escrevendo("Programando em Golang")
}

func escrevendo(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
