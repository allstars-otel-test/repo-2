package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	addr := ":8080"

	log.Printf("Server starting on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}