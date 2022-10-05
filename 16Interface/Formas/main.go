package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

func escrevendoArea(f forma) {
	fmt.Printf("A area da forma é de %0.2f \n", f.area())
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

func main() {
	fmt.Println("Interfaces Formas")

	r := retangulo{10, 15}
	escrevendoArea(r)

	c := circulo{10}
	escrevendoArea(c)
}
