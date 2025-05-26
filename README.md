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



## Development Helpers

The repository includes a `Makefile` and a `docker-compose.yaml` file to set up a local MongoDB instance with some seed data.

```bash
# Build and run MongoDB with sample data
make init-db
```

This uses the official MongoDB Docker image through Docker Compose and executes `scripts/init-db.js` on startup. The `install-mongo` target demonstrates how MongoDB can be installed on Debian/Ubuntu based systems.

