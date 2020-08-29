package main

import (
	"fmt"
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	log.Println("health accessed")
	fmt.Fprintf(w, "{health:ok}")
}

func main() {
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)
}
