package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (u user) save() {
	fmt.Println("Salvando...")
}

func main() {
	person := user{"Tayler", 28}

	user.save(person)
}
