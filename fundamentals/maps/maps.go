// map estrutura para guardar dados

package main

import "fmt"

func main() {
	user := map[string]string{
		"name":     "Marcos",
		"lastname": "Lobato",
	}

	fmt.Println(user)

	// map aninhado
	student := map[string]map[string]string{
		"nome": {
			"primeiro": "joão",
			"ultimo":   "pedro",
		},
		"curso": {
			"Nome":   "Engenharia",
			"Campus": "ZH",
		},
	}

	// remover chave
	delete(student, "nome") // map e chave

	fmt.Println(student)
}
