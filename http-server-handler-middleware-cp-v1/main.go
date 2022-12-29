package main

import (
	"net/http"
)

func RequestMethodGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.WriteHeader(200)
			w.Write([]byte("Welcome to Student page"))
		} else {
			w.WriteHeader(405)
			w.Write([]byte("Method is not allowed"))
		}
	}) // TODO: replace this
}
func StudentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Student page"))
	}
}

func main() {
	http.Handle("/student", RequestMethodGet(StudentHandler()))
	// TODO: answer here
	route := "localhost:8080"
	http.ListenAndServe(route, nil)
}
