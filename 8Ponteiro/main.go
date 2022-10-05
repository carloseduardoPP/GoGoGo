package main

import "fmt"

func main() {
	fmt.Println("Ponteiro")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	//PONTEIRO É UMA REFERENCIA DE MEMORIA
	var variavel3 int
	var Ponteiro *int

	variavel3 = 100
	Ponteiro = &variavel3
	fmt.Println(variavel3, *Ponteiro)
}
