package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Wait Group")

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		escrevendo("Olá mundo")
		waitGroup.Done()
	}()
	go func() {
		escrevendo("Programando em go")
		waitGroup.Done()
	}()
	go func() {
		escrevendo("Goroutines 3")
		waitGroup.Done()
	}()
	go func() {
		escrevendo("Goroutines 4")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrevendo(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
