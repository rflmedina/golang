package main

import "fmt"

func recoverTheCode() {
	if x := recover(); x != nil {
		fmt.Println("Execução recuperada")
	}
}

func media(grade1, grade2 int) string {
	defer recoverTheCode()

	media := (grade1 + grade2) / 2

	if media > 6 {
		fmt.Println("Aluno aprovado")
	} else if media < 6 {
		fmt.Println("Aluno reprovado")
	}

	panic("")
}

func main() {

}
