package main

import "fmt"

var n int

func init() {
	fmt.Println("Executa primeiro que a func main")
	n = 20
}

func main() {
	fmt.Println("executa tudo que por aqui")
	fmt.Println(n)
}
