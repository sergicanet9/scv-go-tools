# scv-go-tools v3

Tools for building REST APIs in Go.

## Included packages
- API/middlewares: provides Middlewares for panic recovering and JWT authentication & role-based authorization.
- API/utils: provides JSON success/error responses with logs.
- INFRASTRUCTURE: provides MongoDB and PostgreSQL connection functions and a generic implemention of the repository interface for MongoDB.
- REPOSITORY: provides an interface of the repository pattern with CRUD operations, to be used in the Core of the API.

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
