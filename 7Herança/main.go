package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint16
	altura    uint16
}

type estudante struct {
	pessoa
	curso     string
	turma     uint16
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Caetano", "Veloso", 25, 170}
	fmt.Println(p1)

	e1 := estudante{p1, "pedagogia", 3, "Uninove"}
	fmt.Println(e1)

}
