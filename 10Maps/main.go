package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro":  "Carlos",
			"sobrenome": "Eduardo",
		},
		"curso": {
			"nome do curso": "engenharia",
			"campus":        "Campus 1",
		},
	}

	fmt.Println(usuario2)
}
