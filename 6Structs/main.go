package main

import "fmt"

type usuarios struct {
	nome  string
	idade uint8
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuarios
	u.nome = "Davi"
	u.idade = 21
	u.numero = 0
	u.logadouro = "rua ciro gomes"
	fmt.Println(u)

	u2 := usuarios{"Davi", 21}
	fmt.Println(u2)

	u3 := usuarios{idade: 21}
	fmt.Println(u3)

}
