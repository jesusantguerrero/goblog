package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world 2")
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is an index page")
}

func main() {

	http.HandleFunc("/", Home)

	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)

	server := &http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}
