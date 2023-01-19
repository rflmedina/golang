package main

import "fmt"

func main() {
	var x int = 10
	var y int = x
	// fmt.Print(y)
	y--
	x++
	// fmt.Print(x)

	// PONTEIRO: referência de memória
	var z int
	var ponteiro *int // endereço de memória de um valor inteiro

	z = 314
	ponteiro = &z

	fmt.Println(z, *ponteiro) // desreferenciação
}
