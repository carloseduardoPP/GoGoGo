package main

import (
	"crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println()

	canal := multiplexiar(escrever("ola mundo"), escrever("Programando em go"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

func multiplexiar(canaldeEntrada1, canaldeEntrada2 <-chan string) <-chan string {
	canaldeSaida := make(chan string)

	go func() {
		select {
		case mensagem := <-canaldeEntrada1:
			canaldeSaida <- mensagem

		case mensagem := <-canaldeEntrada2:
			canaldeSaida <- mensagem

		}
	}()
	return canaldeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()
	return canal
}
