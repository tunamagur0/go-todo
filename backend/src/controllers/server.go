package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tunamagur0/go-todo/models"
	"gorm.io/gorm"
)

type Server struct {
	server *http.Server
	db     *gorm.DB
}

func NewServer(addr string, db *gorm.DB) *Server {
	return &Server{
		server: &http.Server{Addr: addr},
		db:     db,
	}
}

func (s *Server) Start() error {
	s.initHandlers()
	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.server.Shutdown(ctx)
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *Server) HandleHealth(w http.ResponseWriter, r *http.Request) {
	log.Println("health accessed")
	fmt.Fprintf(w, "{health:ok}")
}

func (s *Server) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if todo.Content == "" {
		http.Error(w, "content is empty", http.StatusBadRequest)
		return
	}
	s.db.Create(&todo)
}

func (s *Server) initHandlers() {
	mux := http.NewServeMux()
	s.server.Handler = mux

	mux.HandleFunc("/health", s.HandleHealth)
	mux.HandleFunc("/create", s.HandleCreate)
}
