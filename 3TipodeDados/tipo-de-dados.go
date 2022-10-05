package main

import (
	"errors"
	"fmt"
)

func main() {
	//int8, int16, int32, int64

	var numero int64 = 1000000
	fmt.Println(numero)

	var numero2 uint32 = 10
	fmt.Println(numero2)

	//alias
	//int32 = rune
	var numero3 rune = 309
	fmt.Println(numero3)

	//byte = uint8
	var numero4 byte = 9
	fmt.Println(numero4)

	var numero5 float32 = 400.00
	fmt.Println(numero5)

	var numero6 float64 = 400000.00000
	fmt.Println(numero6)

	// string

	var str string = "Ola Mundo"
	fmt.Println(str)

	str2 := "Ola Céu"
	fmt.Println(str2)

	//tabela ascii
	char := 'B'
	fmt.Println(char)

	//Valor 0 fim string
	var texto string
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("erro interno")
	fmt.Println(erro)
}
