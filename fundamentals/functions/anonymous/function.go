package main

import "fmt"

func main() {

	func(message string) {
		fmt.Println(message)
	}("Que medo alegre, o de te esperar.") // Função anônima

	getAuthor := func() string {
		return "Clarice Lispector"
	}()

	fmt.Println(getAuthor)
}
