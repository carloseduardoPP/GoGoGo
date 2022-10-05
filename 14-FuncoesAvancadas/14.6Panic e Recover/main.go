package main

import "fmt"

func media(n1, n2 float64) bool {
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A media é exatamente 6!")
}

func main() {
	fmt.Println("Panic e Recover")

	fmt.Println(media(6, 7))
	fmt.Println("Pós execução")
}
