package server

import (
	"crud/database"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID    uint32 `json:"id"`
	name  string `json:"name"`
	email string `json:"email"`
}

// CreateUser is a handler function for the /users endpoint
func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Error reading request body"))
		return
	}

	var user user // user is a struct
	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		w.Write([]byte("Error unmarshalling request body"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}

	defer db.Close()

	// Prepare the statement to insert a new user

	statement, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	if err != nil {
		w.Write([]byte("Error preparing statement"))
		return
	}

	defer statement.Close()

	inserction, err := statement.Exec(user.name, user.email)
	if err != nil {
		w.Write([]byte("Error inserting new user"))
		return
	}

	id, err := inserction.LastInsertId()
	if err != nil {
		w.Write([]byte("Error getting last insert id"))
		return
	}

	// Statis code 201 means that a new resource was created
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User successfully created! ID: " + string(id)))
}
