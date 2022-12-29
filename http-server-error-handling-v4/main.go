package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func MethodGet(r *http.Request) error {
	//checking get
	if r.Method != http.MethodGet {
		return fmt.Errorf("Method not allowed")
	}
	return nil
}

func CheckDataRequest(r *http.Request) error {
	//check len data
	dataURL := r.URL.Query().Get("data")
	if len(dataURL) == 0 {
		return fmt.Errorf("Data not found")
	}
	return nil
}

func CheckOpenFile(r *http.Request) error {
	f := r.URL.Query().Get("filename")
	_, err := os.Open(f)
	if err != nil {
		return fmt.Errorf("File not found")
	}
	return nil
}

func MethodHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := MethodGet(r)
		log.Println(err)
		if r.Method == "GET" {
			w.WriteHeader(200)
			w.Write([]byte("Method handler passed"))
		} else {
			w.WriteHeader(405)
			w.Write([]byte("Method not allowed"))
		}
	}
} // TODO: replace this

func DataHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := CheckDataRequest(r)
		fmt.Println(err)
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte("Data not found"))
		} else {
			w.WriteHeader(200)
			w.Write([]byte("Data handler passed"))
		}
		// TODO: replace this
	}
}

func OpenFileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//open
		err := CheckOpenFile(r)
		fmt.Println(err) // TODO: replace this
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("File not found"))
		} else {
			w.WriteHeader(200)
			w.Write([]byte("Error handler passed"))
		}
	}
}
