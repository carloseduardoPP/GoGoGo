package main

import "fmt"

func main() {
	// ARITIMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 4 / 2
	multiplicacao := 2 * 2
	restoDaDivisão := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisão)

	var numero1 int16 = 10
	var numero2 int16 = 15
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//ATRIBUIÇÃO
	var variavel1 string = "string"
	variavel2 := "string 2"
	fmt.Println(variavel1, variavel2)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2, 1 < 2, 1 == 2, 1 >= 2, 1 <= 2, 1 != 2)

	//OPERADORES LOGICOS
	fmt.Println("-----------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//OPERADORES UNARIOS
	numero := 10

	numero++
	numero += 17

	numero--
	numero -= 10

	numero /= 2

	numero *= 3

	numero %= 5
	fmt.Println(numero)

	fmt.Println("---------------------")
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)

}
