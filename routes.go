package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string           `json:"name"`
	Method      string           `json:"method"`
	Path        string           `json:"path"`
	HandlerFunc http.HandlerFunc `json:"-"`
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		// Apply middleware
		handler = Header(handler)
		handler = Logger(handler, route.Name)
		// Register route
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

func GetPath(name string) *string {
	if path, err := AppRouter.Get(name).GetPathTemplate(); err == nil {
		return &path
	} else {
		panic(err)
	}
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"CreateUser",
		"POST",
		"/user",
		CreateUser,
	},
	Route{
		"GetUser",
		"GET",
		"/user/{id}",
		GetUser,
	},
	Route{
		"Auth",
		"POST",
		"/auth",
		Auth,
	},
}
