package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	fmt.Println("Generico")

	generica("String")
	generica(1)
	generica(true)

	// no tipo generico, pode se printar, validar tudo, qualquer coisa, o que acaba sendo um tanto quanto bagunçado
}
