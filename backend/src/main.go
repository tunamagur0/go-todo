package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tunamagur0/go-todo/todo"
)

func health(w http.ResponseWriter, r *http.Request) {
	log.Println("health accessed")
	fmt.Fprintf(w, "{health:ok}")
}

func create(w http.ResponseWriter, r *http.Request) {
	var todo todo.TODO
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if todo.Content == "" {
		http.Error(w, "content is empty", http.StatusBadRequest)
		return
	}

	log.Println(todo.Content)
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/create", create)
	http.ListenAndServe(":8080", nil)
}
