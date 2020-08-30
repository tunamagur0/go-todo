package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

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

func (s *Server) HandleTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only get is allowed", http.StatusMethodNotAllowed)
		return
	}

	var todos []models.Todo
	if err := s.db.Find(&todos).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if todos == nil {
		todos = []models.Todo{}
	}

	res, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (s *Server) HandleTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only get is allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	var todo models.Todo
	query := s.db.Where("id = ?", id).First(&todo)
	if err := query.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (s *Server) HandleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only post is allowed", http.StatusMethodNotAllowed)
		return
	}

	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if todo.Content == "" {
		http.Error(w, "content is empty", http.StatusBadRequest)
		return
	}

	if utf8.RuneCountInString(todo.Content) > 255 {
		http.Error(w, "content is too long", http.StatusBadRequest)
		return
	}

	id, err := uuid.NewRandom()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	todo.ID = id.String()
	s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&todo).Error; err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return err
		}
		return nil
	})

	w.WriteHeader(http.StatusCreated)
	return
}

func (s *Server) initHandlers() {
	r := mux.NewRouter()
	s.server.Handler = r

	r.HandleFunc("/health", s.HandleHealth)
	r.HandleFunc("/todos", s.HandleTodos)
	r.HandleFunc("/create", s.HandleCreate)
	r.HandleFunc("/todo/{id}", s.HandleTodo)
}
