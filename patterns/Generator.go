package main

import (
	"fmt"
	"time"
)

func main() {
	c := write("Hello World")

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func write(message string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- fmt.Sprintln(message)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return c
}
