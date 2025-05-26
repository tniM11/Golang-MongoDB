package mongo

import (
	"context"

	"github.com/example/todos/internal/domain"
	"github.com/example/todos/internal/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TodoRepository is a MongoDB implementation of ports.TodoRepository.
type TodoRepository struct {
	collection *mongo.Collection
}

// NewTodoRepository creates a new Mongo backed repository.
func NewTodoRepository(db *mongo.Database) ports.TodoRepository {
	return &TodoRepository{collection: db.Collection("todos")}
}

func (r *TodoRepository) Create(todo *domain.Todo) error {
	_, err := r.collection.InsertOne(context.Background(), todo)
	return err
}

func (r *TodoRepository) List() ([]domain.Todo, error) {
	cur, err := r.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var todos []domain.Todo
	for cur.Next(context.Background()) {
		var t domain.Todo
		if err := cur.Decode(&t); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, cur.Err()
}
