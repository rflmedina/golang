package main

import "fmt"

// funções são consideradas tipos | tipo func()
// tipos func() podem incluir o tipo primitido de um parâmetro e de um retorno
func sum(x int8, y int8) int8 {
	return x + y
}

// funções no go podem ter mais de um return
func calculator(x, y int8) (int8, int8) {
	sum := x + y
	subs := x - y

	return sum, subs
}

func main() {
	resultado := sum(5, 4)
	fmt.Print(resultado)

	var f = func() string {
		const txt = ""
		return txt
	}

	f()

	// _ não add o valor
	resultado1, _ := calculator(10, 40)
	fmt.Print(resultado1)
}
