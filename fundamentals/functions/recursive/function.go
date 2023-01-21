package main

import "fmt"

func fibonacci(position int) int {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {
	position := fibonacci(4)

	fmt.Print(position)
}
