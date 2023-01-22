package main

import "fmt"

func closure() func() {

	x := func() {
		fmt.Println("Test")
	}

	return x

}

// Nada a comentar really...

func main() {
	// text := "Hey"

	closure()
}
