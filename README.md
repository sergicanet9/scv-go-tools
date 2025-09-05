# scv-go-tools v4
![CI](https://github.com/sergicanet9/scv-go-tools/actions/workflows/ci.yml/badge.svg)
![Coverage](https://img.shields.io/badge/Coverage-97.4%25-brightgreen)
[![Go Reference](https://pkg.go.dev/badge/github.com/sergicanet9/scv-go-tools/v4.svg)](https://pkg.go.dev/github.com/sergicanet9/scv-go-tools/v4)

Toolkit for building REST and gRPC APIs in Go, structured around clean architecture principles.

## Included packages
- **api/middlewares**: HTTP middlewares for panic recovery, JWT authentication, role-based authorization, and request/response logging.
- **api/interceptors**: gRPC interceptors providing equivalent functionality to HTTP middlewares, supporting both unary and stream RPCs.
- **api/utils**: Utility functions for sending HTTP and gRPC success/error responses to clients with proper status code management, and JSON unmarshalling from files with support for parsing `time.Duration`.
- **infrastructure**: Connection management for MongoDB and PostgreSQL, a PostgreSQL migration runner, and a generic MongoDB repository implementation.
- **mocks**: Mock creation for MongoDB and PostgreSQL repositories to facilitate unit testing.
- **observability**: New Relic integration for APM and log forwarding, including a singleton logger.
- **repository**: Interface for the Repository pattern defining CRUD operations, designed for multiple storage implementations and extensibility through composition.
- **wrappers**: Custom type wrappers including specialized error types for simpler error codes mapping and a gRPC Server Stream wrapper for enabling context injection.
- **testutils**: Convinient utility functions to simplify testing.

## Installation
```
go get github.com/sergicanet9/scv-go-tools/v4
```

## Run all unit tests with code coverage
```
make test-unit
```

## View coverage report
```
make cover
```

## Usage examples
Check out [go-hexagonal-api](https://github.com/sergicanet9/go-hexagonal-api) for practical examples of how to use the library with both HTTP and gRPC APIs.

## Author
Sergi Canet Vela

## License
This project is licensed under the terms of the MIT license.
