# scv-go-tools v3
![CI](https://github.com/sergicanet9/scv-go-tools/actions/workflows/pipeline.yml/badge.svg)
![Coverage](https://img.shields.io/badge/Coverage-91.1%25-brightgreen)

Tools for building REST APIs in Go.

## Included packages
- api/middlewares: provides Middlewares for panic recovering and JWT authentication & role-based authorization.
- api/utils: provides JSON success/error responses with logs.
- infrastructure: provides MongoDB and PostgreSQL connection functions and a generic implemention of the Repository interface for MongoDB.
- mocks: provides mock creation functions for MongoDB and PostgreSQL.
- repository: provides an interface of the repository pattern with CRUD operations.

## Usage steps
1. Create an empty repository and clone it.
2. Execute:
```
go mod init github.com/{username}/{repository_name}
go get github.com/sergicanet9/scv-go-tools/v3
```

## Run all unit tests with coverage
```
go test -race ./... -coverprofile=coverage.out
```

## View code coverage
```
go tool cover -html=coverage.out
```

## Usage examples
[go-hexagonal-restapi](https://github.com/sergicanet9/go-hexagonal-api)

## Author
Sergi Canet Vela

## License
This project is licensed under the terms of the MIT license.
