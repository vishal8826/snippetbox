package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting the server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}