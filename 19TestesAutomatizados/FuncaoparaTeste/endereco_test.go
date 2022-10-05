package endereco_test

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenarioDeTeste := []cenarioDeTeste{
		{"Avenida ABC", "Avenida"},
		{"Rua ABC", "Rua"},
		{"Rodovia ABC", "Rodovia"},
		{"Praça ABC", "Praças"},
		{"Estrada ABC", "Estrada"},
		{"", "Tipo Invalido"},
	}

	for _, cenario := range cenarioDeTeste {
		enderecoRecebido := tiposdeEndereco(cenario.enderecoInserido)
		if enderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				enderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

}
