package main

import "fmt"

func Closure() func() {
	text := "dentro da função"

	funcao := func() {
		fmt.Println(text)
	}

	return funcao
}

func main() {
	fmt.Println("Closure")

	texto := "dentro da função main"
	fmt.Println(texto)

	funcaoNova := Closure()
	funcaoNova()
}
