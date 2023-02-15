package main

import (
	"fmt"
	"test/address"
)

func main() {
	addressType := address.AddressType("nice")
	fmt.Println(addressType)
}
