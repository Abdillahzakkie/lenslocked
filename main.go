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

	fmt.Println("Server is running on port:", port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf("localhost:%s", port), router))
}