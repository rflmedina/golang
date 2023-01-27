package main

import "fmt"

func init() {
	fmt.Println("Init")
	// cada arquivo pode ter uma função init
}

func main() {
	fmt.Println("Main")
}
