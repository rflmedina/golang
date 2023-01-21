package main

import "fmt"

func main() {
	x := 10

	if x > 10 { // Não é necessário usar parenteses
		fmt.Println("Maior que 10.")
	} else {
		fmt.Println("Menor que 10")
	}

	if y := 1; y > 0 { // declarnado a variavel e testando se ela atende uma condição
		fmt.Println("Sim")
	}
}
