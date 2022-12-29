package main

import (
	"net/http"
)

func AdminHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Admin page"))
	}
}

func RequestMethodGetMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(405)
			w.Write([]byte("Method is not allowed"))
		} else {
			next.ServeHTTP(w, r)
		}
	}) // TODO: replace this
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("role") != "ADMIN" {
			w.WriteHeader(401)
			w.Write([]byte("Role not authorized"))
		} else {
			next.ServeHTTP(w, r)
		}
	}) // TODO: replace this
}

func main() {
	// TODO: answer here
	mux := http.DefaultServeMux
	mux.HandleFunc("/admin", AdminHandler())
	// http.HandleFunc("/admin", )
	http.ListenAndServe("localhost:8080", nil)
}
