package main

import "fmt"

func main() {
	// ARITIMÉTICOS
	sum := 1 + 2
	sub := 1 - 2
	divider := 10 / 4
	mult := 10 * 2
	product := 10 % 2

	fmt.Print(sum, sub, divider, mult, product)
	// Não é possível fazer operações com tipos numéricos diferentes

	// ATRIBUIÇÃo
	var hello string = "Olá"
	hi := "Oi"
	fmt.Print(hello, hi)

	// RELACIONAIS
	fmt.Print(1 > 2)  // maior que
	fmt.Print(1 >= 2) // maior ou igual a
	fmt.Print(1 == 2) // igual a
	fmt.Print(1 <= 2) // menor ou igual a
	fmt.Print(1 < 2)  // menor que
	fmt.Print(1 != 2) // diferente de

	// LÓGICOS
	fmt.Print(true && true) // se todos forem true
	fmt.Print(true || true) // caso se um for true
	fmt.Print(!true)        // inverter o valor de uma variavél

	// UNÁRIOS
	x := 10
	x++
	x += 20
	x--
	x -= 20

	// Não existe operador ternário em GO.
}
