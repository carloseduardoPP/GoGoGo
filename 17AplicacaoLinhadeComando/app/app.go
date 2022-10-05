package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli" 
)

//Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.name = "Aplicação de LInha de Comando"
	app.Usage = "Busca IP's e nomes de servidores na rede"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app.Commandans = []cli.Command{
		{
			Name:   "IP",
			Usage:  "Busca Ip's de endereços na rede",
			Flags:  flags,
			Action: buscarIps,
		},
	}
	{
		Name: "servidores",
		Usage: "Buscar o nome do servidores na internet",
		Flags: flags,
		Action: buscarServidores,
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servidores, erro := net.LookupNS(host) 
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
