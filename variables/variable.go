package main

import "fmt"

func main() {
	var nome string = "Agatha" // declarando de modo explícito
	sobrenome := " Cristhie"   // declarando de modo implícito - inferência de tipo
	fmt.Println(nome + sobrenome)

	var (
		escolaridade string = ""
		hobbie       string = ""
	)

	fmt.Println(escolaridade, hobbie)

	x, y := 10, 50

	fmt.Println(x + y)

	const umaConst string = "teste"
	fmt.Println(umaConst)

}
