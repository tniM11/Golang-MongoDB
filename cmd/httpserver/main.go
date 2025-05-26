package main

import (
	"context"
	"log"
	"net/http"
	"time"

	adapter_mdb "github.com/example/todos/internal/adapters/mongo"
	httpHandler "github.com/example/todos/internal/handlers/http"
	"github.com/example/todos/internal/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatalf("cannot connect to mongo: %v", err)
	}
	db := client.Database("todos")

	repo := adapter_mdb.NewTodoRepository(db)
	service := services.NewTodoService(repo)
	handler := httpHandler.NewHandler(service)

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handler.CreateTodo(w, r)
		case http.MethodGet:
			handler.ListTodos(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	log.Println("server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
