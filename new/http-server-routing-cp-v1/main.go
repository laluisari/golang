package main

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%s, %v %s %v", time.Now().Weekday(), time.Now().Day(), time.Now().Month(), time.Now().Year())))
	}) // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name != "" {
			w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
		} else {
			w.Write([]byte("Hello there"))
		}
	}) // TODO: replace this
}

func main() {
	// TODO: answer here
	http.ListenAndServe("localhost:8080", nil)
}
