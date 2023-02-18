package server

import (
	"crud/database"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// SearchUsers is a handler function for the /users endpoint
func SearchUsers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}

	defer db.Close()

	// Prepare the statement to select all users
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		w.Write([]byte("Error at getting users"))
		return
	}

	defer rows.Close()

	var users []user

	for rows.Next() {
		var user user

		if err = rows.Scan(&user.ID, &user.name, &user.email); err != nil {
			w.Write([]byte("Error scanning users"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Error encoding users"))
		return
	}
}

// SearchUser is a handler function for the /users/{id} endpoint
func SearchUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Error converting id to uint"))
		return
	}

	db, err := database.Connect()
	if err != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}

	defer db.Close()

	// Prepare the statement to select a user by id
	row, err := db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		w.Write([]byte("Error at getting user"))
		return
	}

	defer row.Close()

	var user user
	if row.Next() {
		if err = row.Scan(&user.ID, &user.name, &user.email); err != nil {
			w.Write([]byte("Error scanning user"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Error encoding user"))
		return
	}

}
