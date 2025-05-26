package services

import "github.com/example/todos/internal/domain"
import "github.com/example/todos/internal/ports"

// TodoService provides todo related operations.
type TodoService struct {
	repo ports.TodoRepository
}

func NewTodoService(r ports.TodoRepository) *TodoService {
	return &TodoService{repo: r}
}

func (s *TodoService) Create(todo *domain.Todo) error {
	return s.repo.Create(todo)
}

func (s *TodoService) List() ([]domain.Todo, error) {
	return s.repo.List()
}
