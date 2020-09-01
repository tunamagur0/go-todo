package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/tunamagur0/go-todo/controllers"
	"github.com/tunamagur0/go-todo/db"
)

func run(addr string) int {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	errCh := make(chan error)
	db := db.Init()
	server := controllers.NewServer(addr, db)

	go func() {
		errCh <- server.Start()
	}()

	select {
	case err := <-errCh:
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
	case <-sigCh:
		if err := server.Stop(context.Background()); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
	}

	return 0
}

func main() {
	port, _ := strconv.Atoi(os.Args[1])
	os.Exit(run(fmt.Sprintf(":%d", port)))
}
