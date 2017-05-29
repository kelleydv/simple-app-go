package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	statusSuccess = "success"
	statusError   = "error"
)

type jsonResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type jsonResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func jsendResponse(w *http.ResponseWriter, payload interface{}, errorMessage *string) {
	if errorMessage == nil {
		resp := jsonResponse{statusSuccess, payload}
		json.NewEncoder(*w).Encode(&resp)
	} else {
		resp := jsonResponseError{statusError, *errorMessage}
		json.NewEncoder(*w).Encode(&resp)
	}
}

var AppRouter *mux.Router //needs to be set

// Render the main interface
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	// Get application paths to pass to interface
	data := struct {
		LoginPath    *string
		RegisterPath *string
	}{
		GetPath("Auth"),
		GetPath("CreateUser"),
	}
	t, _ := template.ParseFiles("static/index.html")
	t.Execute(w, &data)
}

// Save a user record
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user = new(User)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		errmsg := err.Error()
		http.Error(w, errmsg, http.StatusBadRequest)
		jsendResponse(&w, nil, &errmsg)
		return
	}
	if user.Username == "" || user.Password == "" {
		errmsg := "missing username or password"
		http.Error(w, errmsg, http.StatusBadRequest)
		jsendResponse(&w, nil, &errmsg)
		return
	}
	if err := user.Create(); err != nil {
		errmsg := err.Error()
		http.Error(w, errmsg, http.StatusConflict)
		jsendResponse(&w, nil, &errmsg)
		return
	}
	jsendResponse(&w, user, nil)
}

// Get a user record from an id
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user = new(User)
	err := user.Get(&id)
	if err != nil {
		errmsg := err.Error()
		http.Error(w, errmsg, http.StatusBadRequest)
		jsendResponse(&w, nil, &errmsg)
		return
	}
	jsendResponse(&w, user, nil)
}

// Authenticate with a username and password
func Auth(w http.ResponseWriter, r *http.Request) {
	var user = new(User)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		errmsg := err.Error()
		http.Error(w, errmsg, http.StatusBadRequest)
		jsendResponse(&w, nil, &errmsg)
		return
	}
	if result, err := user.Auth(); !result {
		errmsg := "failed auth"
		if err != nil {
			errStr := fmt.Sprintf("%s %s", errmsg, err.Error())
			http.Error(w, errStr, http.StatusUnauthorized)
			jsendResponse(&w, nil, &errStr)
			return
		}
		http.Error(w, errmsg, http.StatusUnauthorized)
		jsendResponse(&w, nil, &errmsg)
		return
	}
	jsendResponse(&w, user, nil)
}
