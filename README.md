# gRPC Gateway Service

This project implements a simple greeting service using gRPC and gRPC-Gateway. It demonstrates how to create a service that can handle both gRPC calls and REST API requests.

## Prerequisites

Before you begin, ensure you have the following installed:
- Go (1.16 or later)
- Protocol Buffers Compiler (protoc)

Required Go packages:
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
```

## Project Structure

```
.
├── api/
│   ├── greeting.proto
│   ├── greeting.pb.go       # Generated
│   ├── greeting_grpc.pb.go  # Generated
│   └── greeting_pb.gw.go    # Generated
├── server/
│   ├── gateway.go
│   └── server.go
├── go.mod
└── go.sum
```

## Installation

1. Clone the repository:
```bash
git clone <https://github.com/Monish-2304/Go-Grpc-Gateway>
cd <project-directory>
```

2. Install dependencies:
```bash
go mod tidy
```

3. Generate Proto files:
```powershell
# For Windows PowerShell
protoc -I. `
    -I ${env:GOPATH}\pkg\mod\github.com\grpc-ecosystem\grpc-gateway@v2 `
    -I ${env:GOPATH}\pkg\mod\github.com\googleapis\googleapis@v0.0.0-20241114014203-6b5d85c66e08 `
    --go_out=. `
    --go-grpc_out=. `
    --grpc-gateway_out=. `
    api\greeting.proto

# For Linux/Mac
protoc -I. \
    -I ${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v2 \
    -I ${GOPATH}/pkg/mod/github.com/googleapis/googleapis@v0.0.0-20241114014203-6b5d85c66e08 \
    --go_out=. \
    --go-grpc_out=. \
    --grpc-gateway_out=. \
    api/greeting.proto
```

## Running the Service

1. Start the server:
   You need to start both gateway and server simultaneously
```bash
go run server/gateway.go server/server.go
```

The service will start:
- gRPC server on port 50051
- HTTP gateway on port 8080

## Testing the API

### Using Postman

1. Send a POST request:
- URL: `http://localhost:8080/v1/sayhello`
- Method: `POST`
- Headers: 
  - `Content-Type`: `application/json`
- Body:
```json
{
    "name": "John"
}
```

### Using cURL

```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"John"}' http://localhost:8080/v1/sayhello
```

Expected Response:
```json
{
    "message": "Hello, John!"
}
```



### Rebuild Proto Files

If you modify the proto file, regenerate the Go code using the protoc command mentioned in the Installation section.

## Troubleshooting

1. If you get "Failed to extract ServerMetadata from context":
   - Make sure both gRPC server and gateway are running
   - Check if ports 50051 and 8080 are available

2. If proto generation fails:
   - Verify all required protoc plugins are installed
   - Check GOPATH is correctly set
   - Ensure all required imports are available
