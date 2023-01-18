package main

import (
	"fmt"
	"modulo/helpers"

	"github.com/badoux/checkmail"
)

// Write:  registra uma mensagem na tela.
func main() {
	helpers.Write()

	validateThisEmail := checkmail.ValidateFormat("myemail")
	fmt.Println(validateThisEmail)
}
