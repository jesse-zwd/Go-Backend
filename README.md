# A backend in Go

This is a backend in Go using Gin, Postgresql.

## Core packages

1. Gin - Web framework in Go
2. Gorm - ORM
3. Postgresql - Database
4. Gin-Swagger - Doc
5. Go-Redis - Cache

## Features

1. user/Login - login
2. user/Register - register
3. user/Logout - logout
4. user/me - user information


## Run it locally

For more information of gin-swagger, please check github.com/swaggo/gin-swagger.

### create database on Postgresql
### set configuration in .env file 
### go mod tidy
### swag init
### go run main.go

On a browser, open http://localhost:9000/swagger/index.html, you can test the APIs.

## License

MIT