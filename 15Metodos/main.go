package main

import "fmt"

type Usuario struct {
	nome  string
	idade uint8
}

func (u Usuario) salvar() {
	fmt.Printf("Salvando os Dados do %s no Banco de Dados do Sistema \n", u.nome)
}

func (u Usuario) maiorIdade() bool {
	return u.idade >= 18
}

func (u *Usuario) fazerAniver() {
	u.idade++
}

func main() {
	Usuario1 := Usuario{"joão", 30}
	Usuario1.salvar()

	Usuario2 := Usuario{"Davi", 50}
	Usuario2.salvar()

	maiorIdade := Usuario2.maiorIdade()
	fmt.Println(maiorIdade)

	Usuario2.fazerAniver()
	fmt.Println(Usuario2.idade)
}
