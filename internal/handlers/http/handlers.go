package http

import (
	"encoding/json"
	"net/http"

	"github.com/example/todos/internal/domain"
	"github.com/example/todos/internal/services"
)

// Handler provides HTTP endpoints for todos.
type Handler struct {
	service *services.TodoService
}

func NewHandler(s *services.TodoService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var t domain.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.Create(&t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(t)
}

func (h *Handler) ListTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.NewEncoder(w).Encode(todos)
}
