package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

// Render the main interface
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// Save a user record
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user = new(User)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user.Username == "" || user.Password == "" {
		http.Error(w, "missing username or password", http.StatusBadRequest)
		return
	}
	if err := user.Create(); err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
	}
}

// Get a user record from an id
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user = new(User)
	err := user.Get(&id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(user)
}

// Authenticate with a username and password
func Auth(w http.ResponseWriter, r *http.Request) {
	var user = new(User)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if result, err := user.Auth(); !result {
		if err != nil {
			errStr := fmt.Sprintf("failed auth %v", err.Error())
			http.Error(w, errStr, http.StatusUnauthorized)
		}
		http.Error(w, "failed auth", http.StatusUnauthorized)
		return
	}
}
