package main

import (
	"fmt"
	"log"
	http "net/http"

	mux "github.com/gorilla/mux"
)

func main() {
	port := "8080"
	router := mux.NewRouter()

	// match all unknown routes
	router.NotFoundHandler = http.HandlerFunc(notFound)

	// handle default GET "/" route
	router.HandleFunc("/", Index).Methods(http.MethodGet)

	fmt.Println("Server is running on port:", port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf("localhost:%s", port), router))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	http.Error(w, "404 - Page Not Found", http.StatusNotFound)
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf8")
	fmt.Fprintln(w, "Welcome to Home page...")
}