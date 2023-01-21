package main

import "fmt"

// Só existe uma estrutura de loop: FOR

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	names := [3]string{"João", "Lucas", "Pedro"}

	for _, name := range names {
		fmt.Println(name)
	}
}
