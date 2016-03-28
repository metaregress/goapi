package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Print("Server starting.")
	log.Fatal(http.ListenAndServe(":8081", router))
}
