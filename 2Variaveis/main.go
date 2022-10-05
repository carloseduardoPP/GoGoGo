package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "lala"
		variavel4 string = "lulu"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "555", "666"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	var conta int = 1000
	fmt.Println(conta)

	var numero int = 1
	fmt.Println(numero)

	sum := 116 - 68
	fmt.Println(sum)

}
