package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("file server running on :5050")
	log.Fatal(http.ListenAndServe(":5050", http.FileServer(http.Dir("."))))
}
