# Golang MongoDB Todo App

This project is a simple Todo application written in Go using MongoDB. It demonstrates a minimal hexagonal architecture separating domain, ports, adapters and services.

## Running

The application expects a running MongoDB instance. By default it connects to `mongodb://localhost:27017` and uses database `todos`.

```
go run ./cmd/httpserver
```

This starts an HTTP server on `:8080` with the following endpoints:

- `POST /todos` – create a new todo
- `GET  /todos` – list all todos


