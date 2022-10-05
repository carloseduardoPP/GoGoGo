package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {

	auxiliar.Escrevendo()

	fmt.Println("Escrevendo alguma coisa")
	fmt.Println("Oi, CarlosJvzz")

	checkmail.ValidateFormat("carlos.caetano@picpay.com")
	fmt.Println("erro")
}
