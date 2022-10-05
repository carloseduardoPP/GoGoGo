package main

import "fmt"

func main() {
	fmt.Println("Buffer")

	canal := make(chan string, 2)
	canal <- "olá mundo"

	mensagem := <-canal
	fmt.Println(mensagem)
}
