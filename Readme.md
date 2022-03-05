# go-server-grpc

This is basic server that has been written with go language.This server contol the username and password which are coming from request and after checking the values return Success or Error

APIs of this server are created with grpc

## Server

> For running the server only `go` should be installed. But if you change the `login.proto` file you should regenerate proto file so you should install `protoc` for generating files.For generating you can run `generate.sh` file

> After installed `go` open in command shell in this directory `login/login_server/` and run this command `go run ./server.go`
