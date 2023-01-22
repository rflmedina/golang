package main

import "fmt"

func sayHi() {
	fmt.Println("Hello")
}

func askHowAre() {
	fmt.Println("how are you?")
}

func main() {
	defer sayHi()
	askHowAre()
}
