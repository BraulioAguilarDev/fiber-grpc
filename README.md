# Fiber with gRPC

## Project structure

1. `proto` directory
Defines the gRPC service and messages using Protocol Buffers

2. `server` directory
Implements the gRPC server and includes the Fiber HTTP server

3. `client` directory
Implements a simple gRPC client that connects to the gRPC server and sends a SayHello request.

## Test code

```sh
// Run server
$ go run server/main.go

// Run client
$  go run client/main.go

// Test HTTP endpoint
$ curl http://localhost:3000/hello/developer
```