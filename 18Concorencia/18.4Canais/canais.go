package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Canais")

	canal := make(chan string)
	go escrevendo("Olá Mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		mensagem, aberto := <-canal
		if !aberto {
			break
		}

		fmt.Println(mensagem)
	}

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa")

}

func escrevendo(texto string, canal chan string) {
	//time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto
		//fmt.Println(texto)
		time.Sleep(time.Second)
	}

	close(canal)
}
