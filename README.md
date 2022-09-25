## Go HTTP API - Hexagonal Architecture

This repository contains a simple HTTP API written in Go. It is a simple example of a hexagonal architecture.

### Requirements

- Go v1.18+
- MySQL (see below).

### Contents

- `cmd/` - contains the main application.
- `internal/` - contains the application code.
- `kit/` - contains the code that can be shared across multiple applications.


### Running the application

The application requires a MySQL database. You can run one using Docker:

```bash
docker run --name go-hexagonal-architecture -e MYSQL_ROOT_PASSWORD=secret -e MYSQL_DATABASE=go-hexagonal-architecture -p 3306:3306 -d mysql:5.7
```

Then, you can run the application:

```bash
go run cmd/main.go
```

### Running the tests

```bash
go test ./...
```

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
