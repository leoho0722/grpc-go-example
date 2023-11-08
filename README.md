# grpc-go-example

Example of implementing gRPC using Go

## Developement Environment

* Apple MacBook Pro (13-inch, M1, 2020)
* macOS Sonoma 14.1 (23B74)
* go@1.21.3
* Visual Studio Code

## Install

```shell
brew install go
brew install protobuf
brew install protoc-gen-go-grpc
brew install protoc-gen-go
```

## Run

### Server side

```shell
go run main.go
```

or

```shell
make run-server
```

### Client side

```shell
go run client/client.go
```

or

```shell
make run-client
```
