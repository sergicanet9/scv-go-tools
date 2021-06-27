# scv-go-framework

Base framework for creating REST APIs in Go.

## Included packages:
- INFRASTRUCTURE: provides MongoDB connection function and Repository Pattern with CRUD operations using the official [mongo--driver](https://github.com/mongodb/mongo-go-driver).
- API/UTILS: provides JSON success/error responses and Middlewares for error handling and JWT token-based authorization.

## Usage steps:
1. Create an empty repository and clone it.
2. Execute:
```
go mod init github.com/{username}/{repository_name}
go env -w GOPRIVATE=github.com/scanet9
go get github.com/scanet9/scv-go-framework 
```
3. As this repository is private, it is needed to create an access token [here](https://github.com/settings/tokens). Give it repo and read:packages permissions.

## Usage example:
[go-mongo-restapi](https://github.com/scanet9/go-mongo-restapi)

