package main

import (
	"context"

	"github.com/tunamagur0/go-todo/controllers"
	"github.com/tunamagur0/go-todo/db"
)

func main() {
	db := db.Init()
	server := controllers.NewServer(":8080", db)
	server.Start()
	defer server.Stop(context.Background())
}
