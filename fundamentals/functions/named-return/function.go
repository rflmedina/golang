package main

import "fmt"

// Retorno nomeado
func calculator(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y

	return
}

func main() {
	sum, sub := calculator(10, 20)

	fmt.Println("soma:", sum, "subtração:", sub)
}
