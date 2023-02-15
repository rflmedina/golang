package address

import "strings"

// AddressType verify if the address type is valid
func AddressType(address string) string {
	validType := []string{"rua", "avenida", "estrada", "rodovia"}

	addressLower := strings.ToLower(address)
	firstWord := strings.Split(addressLower, " ")[0]

	haveAVailidType := false

	for _, v := range validType {
		if v == firstWord {
			haveAVailidType = true
		}
	}

	if haveAVailidType {
		return firstWord
	}

	return "Tipo inválido"
}
