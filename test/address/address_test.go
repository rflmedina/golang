package address

import "testing"

func TestAddressTyper(t *testing.T) {
	addressTypeTeste := "rua dos bobos"
	expected := "rua"
	recived := AddressType(addressTypeTeste)

	if recived != expected {
		t.Errorf("Expected %s, but recived %s", expected, recived)
	}

}
