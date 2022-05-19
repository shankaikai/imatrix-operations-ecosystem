# Capstone S20 Trying Stuff Out Backend

## Pre-requisits
install the go mysql library:
```
go get -u github.com/go-sql-driver/mysql
```
for the env file:
```
go get github.com/joho/godotenv
```
Follow the gRPC quick set up for go: https://grpc.io/docs/languages/go/quickstart/

## Protocol Buffer Generation
```
protoc --go_out=backend --go_opt=paths=source_relative --go-grpc_out=backend --go-grpc_opt=paths=source_relative proto/testing.proto
```

## Start Server
```
go run . 
```