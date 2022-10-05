package main

import "fmt"

func um() {
	fmt.Println("escrevendo a funçãos um")
}

func dois() {
	fmt.Println("escrevndo a função dois ")
}

func main() {
	defer um()
	// defer = adiar
	dois()
}
