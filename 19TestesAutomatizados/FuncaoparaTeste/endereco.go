package endereco

import "strings"

func tiposdeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}

	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiraPalavradoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavradoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavradoEndereco)
	}
	return "Tipo Invalido"
}
