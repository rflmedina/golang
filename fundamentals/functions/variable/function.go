package main

import "fmt"

// Função variática: não é necessário especificar a quantidade de parametros
func sum(x ...int) int { // Só pode ter um parâmetro variático por função e o mesmo deve ficar em último
	total := 0

	for _, x := range x {
		total += x
	}

	return total
}

func main() {
	sumTotal := sum(10, 20, 34, 2345, 10)

	fmt.Println(sumTotal)
}
