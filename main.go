package main

import (
	"log"
	"net/http"
)

func main() {
	AppRouter = NewRouter() // declared in handlers.go
	AppRouter.
		PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(":8080", AppRouter))
}
