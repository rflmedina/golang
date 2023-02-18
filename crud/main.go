package main

import (
	"crud/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8000", router))
}
