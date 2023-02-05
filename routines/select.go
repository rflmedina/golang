package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Microsecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Microsecond * 500)
			canal2 <- "Canal 2"
		}
	}()

	for {
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Print(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Print(mensagemCanal2)

		}

	}
}
