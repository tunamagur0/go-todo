package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/tunamagur0/go-todo/controllers"
	"github.com/tunamagur0/go-todo/db"
)

func run(addr string, dbPath string) int {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	errCh := make(chan error)
	db := db.Init(dbPath)
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
	flagAddr := flag.String("addr", ":8080", "host:port")
	flagDBPath := flag.String("db", "todo.db", "db path")
	flag.Parse()
	os.Exit(run(*flagAddr, *flagDBPath))
}
