# scv-go-tools v3

Tools for building REST APIs in Go, easing the Hexagonal Architecture (Ports & Adapters).

## Included packages
- PORTS: provides an interface of the repository pattern with CRUD operations, to be used in the Core of the API.
- INFRASTRUCTURE: provides MongoDB and PostgreSQL connection functions and adapters for implementing the repository interface.
- API/UTILS: provides JSON success/error responses with logs and Middlewares for error handling and JWT authentication & role-based authorization.

## Usage steps
1. Create an empty repository and clone it.
2. Execute:
```
go mod init github.com/{username}/{repository_name}
go get github.com/sergicanet9/scv-go-tools/v3
```

## Usage examples
[go-hexagonal-restapi](https://github.com/sergicanet9/go-hexagonal-api)

## Author
Sergi Canet Vela

## License
This project is licensed under the terms of the MIT license.
