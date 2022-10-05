package main

import "fmt"

func sinal(numero int) int {
	return numero * -1
}

func sinalcomPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := sinal(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	sinalcomPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
