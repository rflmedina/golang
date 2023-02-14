package address_test

import (
	"test/address"
	"testing"
)

type cenarioDeTeste struct {
	addressTypeTeste string
	expectedTest     string
}

func TestAddressTyper(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"rua dos bobos", "rua"},
		{"avenida dos bobos", "avenida"},
		{"estrada dos bobos", "estrada"},
		{"rodovia dos bobos", "rodovia"},
		{"praça dos bobos", "Tipo inválido"},
		{"r. dos bobos", "Tipo inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		addressTypeReceived := address.AddressType(cenario.addressTypeTeste)

		if addressTypeReceived != cenario.expectedTest {
			t.Errorf("O tipo recebido %s é diferente do esperado %s", addressTypeReceived, cenario.expectedTest)
		}
	}

}
