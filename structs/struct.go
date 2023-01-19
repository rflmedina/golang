// Não existe classe no GO, mas o Struct serve para o mesmo fim.

package main

import "fmt"

type user struct {
	name string
	age  int8
}

func main() {
	var person1 user
	person1.name = "Rafaela"
	person1.age = 19

	fmt.Print(person1)

	person2 := user{"Raiden", 24}

	fmt.Print(person2)

	person3 := user{name: "Agatha"}
	fmt.Print(person3)
}
