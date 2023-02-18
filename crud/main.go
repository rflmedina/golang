package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.handleFunc("/users", func(w http.ResponseWriter, r *http.Request)).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8000", router))
}
