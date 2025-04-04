package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Freishmidt wesite!"))
}

func articleView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific article..."))
}

func articleCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new article..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/article/view", articleView)
	mux.HandleFunc("/article/create", articleCreate)

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
