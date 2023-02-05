// Sincronizar go routines

package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go write("Hello", canal)

	message := <-canal // recebendo valor

	fmt.Println(message)

}

func write(text string, canal chan string) {
	for i := 0; i < 10; i++ {
		canal <- text // atribuição
		time.Sleep(time.Second)
	}

	close(canal)
}
