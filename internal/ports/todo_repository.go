package ports

import "github.com/example/todos/internal/domain"

// TodoRepository defines persistence operations for Todo entities.
type TodoRepository interface {
	Create(todo *domain.Todo) error
	List() ([]domain.Todo, error)
}
