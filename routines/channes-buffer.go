package main

import "fmt"

func main() {
	canal := make(chan string, 2) // buffer: só bloqueia quando ele atinge sua capacidade máxima
	canal <- "Olá Mundo"

	mensagem := <-canal
	fmt.Print(mensagem)
}
