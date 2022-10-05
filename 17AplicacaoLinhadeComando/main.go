package main

import (
	"LinhadeComando/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println()

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
