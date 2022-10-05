package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")

	var Arrays1 [5]int
	Arrays1[0] = 1
	fmt.Println(Arrays1)

	Arrays2 := [5]string{"posição 1", "posição 2", "posição 3", "posição 4", "posição 5"}
	fmt.Println(Arrays2)

	Arrays3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(Arrays3)

	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := Arrays2[1:3]
	fmt.Println(slice2)

	Arrays2[1] = "Posição alterada"
	fmt.Println(slice2)

	//Arrays internos
	fmt.Println("-----------------")
	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //leght
	fmt.Println(cap(slice3)) //capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
