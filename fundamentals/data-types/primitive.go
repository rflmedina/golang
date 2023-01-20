package main

import (
	"errors"
	"fmt"
)

func main() {
	// Tipo int: baseia-se na arquitetura do sistema
	x := 1201

	fmt.Println(x)

	//alias: rune === int32 | byte === unit8

	var y byte = 8
	var m rune = 8

	fmt.Println(y, m)

	// FIM INTEIROS

	var pie float32 = 3.14
	fmt.Println(pie)
	// reais funcionam de forma similar aos inteiros

	// FIM REAIS

	var texto string = "oiiee"

	// ' ' -> é considerado um caractere único, retornando sua posição na tabela ASCII

	fmt.Println(texto)
	// FIM STRINGS

	// Valor Zero - valor inicial
	// sem declarar um valor a variavél recebe o valor inicial padrão

	const iwilllearngo bool = true
	fmt.Println(iwilllearngo)

	// tipo error: muito utilizado, não retorna uma string e sim um tipo de error!!
	const tipodeErro error = errors.New("Erro inteiro")
	fmt.Println(tipodeErro)
}
