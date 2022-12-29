package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
)

func Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(405)
			jsonData, _ := json.Marshal(model.ErrorResponse{
				Error: "Method is not allowed!",
			})
			w.Write(jsonData)
			return
		}
		next.ServeHTTP(w, r)
	}) // TODO: replace this
}

func Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(405)
			jsonData, _ := json.Marshal(model.ErrorResponse{
				Error: "Method is not allowed!",
			})
			w.Write(jsonData)
			return
		}
		next.ServeHTTP(w, r)
	}) // TODO: replace this
}

func Delete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			w.WriteHeader(405)
			jsonData, _ := json.Marshal(model.ErrorResponse{
				Error: "Method is not allowed!",
			})
			w.Write(jsonData)
			return
		}
		next.ServeHTTP(w, r)
	}) // TODO: replace this
}
