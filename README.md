## Go HTTP API - Hexagonal Architecture

This repository contains a simple HTTP API written in Go using Gin. It is a simple example of a hexagonal architecture.

Endpoints:
- [ ] GET /health - Health check
- [ ] POST /courses - Create a new course
- [ ] GET /courses - List all courses

Based on the project of the course [Go HTTP API - Hexagonal Architecture](https://pro.codely.com/library/api-http-en-go-aplicando-arquitectura-hexagonal-63367/) by CodelyTV.

### Requirements

- Go v1.18+
- MySQL (see below).

### Contents

- `cmd/` - contains the main application.
- `internal/` - contains the application code.
- `kit/` - contains the code that can be shared across multiple applications.


### Running the application

The application requires a MySQL database. You can run one using connection details in 
'cmd/api/bootstrap/boostrap.go' file and in internal/platform/storage/mysql.

Then, you can run the application:

```bash
go run cmd/api/main.go
```

### Running the tests

```bash
go test ./...
```

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

### TODO

- [ ] Implement EventBus
- [ ] Dockerize the application
- [ ] Upload to Heroku