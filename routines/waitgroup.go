package main

import (
	"fmt"
	"sync"
	"time"
)

func write(message string) {
	for i := 0; i < 10; i++ {
		fmt.Println(message)
		time.Sleep(time.Second)
	}
}

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		write("Hello")
		waitGroup.Done()
	}()

	go func() {
		write("Hello")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}
