package main

import "fmt"

func main() {
	// Todos os elementos dentro de um array devem ser do mesmo tipo.
	// Sempre deve-se declar o tamanho de um array.
	var cores [5]string

	cores[0] = "azul"

	numeros := [...]int{1, 2, 3, 4, 5, 6} // tamanho baseado na quantidade de valores que você passou.

	fmt.Println(cores)
	fmt.Println(numeros)

	// SLICE

	slice := []int{10, 20, 3} // tamanho dinâmico, mas os elementos devem ter o mesmo tipo.

	fmt.Println(slice)

	newSlice := append(slice, 7)

	fmt.Println(newSlice)

	// Arrays Internos
	slice3 := make([]float32, 10, 15) // tipo, quantidade, capacidade.
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length.
	fmt.Println(cap(slice3)) // capacidade.

	// quando seu slice passa de capacidade, ele cria um array interno, dobrando o valor.

}
