package main

import (
	"fmt"
	"time"
)

func main() {
	c := multiplexar(write("Olá Mundo"), write("Programando em Go!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalEntrada1:
				canalSaida <- mensagem
			case mensagem := <-canalEntrada2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
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
