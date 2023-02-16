package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	stringConnection := "golang:golang@/dev?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", stringConnection)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectado ao banco de dados")

	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println(rows)
}
