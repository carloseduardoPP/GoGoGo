package main

import (
	"fmt"
)

func somar(N1 int8, N2 int8) int8 {
	return N1 + N2
}

func calculoMatematico(n1, n2 int8) (int8, int8) {
	somar := n1 + n2
	subtracao := n1 - n2
	return somar, subtracao
}

//////////////////////////////////////////////////////////////

func main() {
	somar := somar(10, 20)
	fmt.Println(somar)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Texto a ser exibido")
	fmt.Println(resultado)

	resultadoSomar, resultadoSubtracao := calculoMatematico(
		10, 15)
	fmt.Println(resultadoSomar, resultadoSubtracao)
}
