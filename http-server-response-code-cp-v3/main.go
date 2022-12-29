package main

import (
	"net/http"
)

var students = []string{
	"Aditira",
	"Dito",
	"Afis",
	"Eddy",
}

func IsNameExists(name string) bool {

	for a := 0; a < len(students); a++ {
		if students[a] == name {
			return true
		}
	}

	return false
}

func CheckStudentName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sname := r.URL.Query().Get("name")
		if r.Method != "GET" {
			w.WriteHeader(405)
			w.Write([]byte("Method is not allowed"))
		} else if IsNameExists(sname) {
			w.WriteHeader(200)
			w.Write([]byte("Name is exists"))
		} else {
			w.WriteHeader(404)
			w.Write([]byte("Data not found"))
		}
	} // TODO: replace this
}

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/students", CheckStudentName())

	return mux
}

func main() {
	route := "localhost:8080"
	http.ListenAndServe(route, GetMux())
}
