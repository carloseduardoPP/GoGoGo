package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando J", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nomes := range nomes {
		fmt.Println(indice, nomes)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	type usuariostruct struct {
		nome      string
		sobrenome string
	}
	usuario2 := usuariostruct{"Zé", "Junior"}

	for chave, valor := range usuario2 {
		fmt.Println(chave, valor)
	}

	for true {
		fmt.Println("Executanto isso infinitamente")
		time.Sleep(time.Second)
	}
}
