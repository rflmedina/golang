package main

import (
	"fmt"
	"time"
)

func write(message string) {
	for {
		fmt.Println(message)
		time.Sleep(time.Second)
	}
}

func main() {
	go write("Hello") // goroutine
	write("----")
}
