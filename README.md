# scv-go-framework v2

Base framework for creating REST APIs in Go.

## Included packages
- INFRASTRUCTURE/MONGO: provides MongoDB connection function and Repository Pattern with CRUD operations using the official [mongo-driver](https://github.com/mongodb/mongo-go-driver).
- INFRASTRUCTURE/POSTGRES: provides PostgreSQL connection function and an interface to implement the Repository Pattern using Go´s included [database/sql package](http://go-database-sql.org).

- API/UTILS: provides JSON success/error responses and Middlewares for error handling and JWT token-based authorization.

## Usage steps
1. Create an empty repository and clone it.
2. Execute:
```
go mod init github.com/{username}/{repository_name}
go get github.com/scanet9/scv-go-framework 
```

## Usage examples
[go-mongo-restapi](https://github.com/sergicanet9/go-mongo-restapi)
<br />
[go-postgres-restapi](https://github.com/sergicanet9/go-postgres-restapi)

## Author
Sergi Canet Vela

## License
This project is licensed under the terms of the MIT license.
