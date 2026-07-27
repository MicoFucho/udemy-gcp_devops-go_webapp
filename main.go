package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	http.HandleFunc("/", requestHandler)
	log.Println("Starting Server...")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	userName := os.Getenv("USERNAME")
	if userName == "" {
		userName = "World"
	}
	fmt.Fprintf(w, "Hello %s!\n", userName)
}
