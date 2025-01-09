# API Gateway with Service Registry (Hybrid HTTP & gRPC)

## Overview

This project implements an API Gateway that routes requests to different microservices using HTTP and gRPC. The API Gateway handles:

HTTP Reverse Proxy for user-service (REST API)

gRPC Proxy for order-service (gRPC API)

JWT Authentication for secure access

Service Registry Removal (static service mapping for simplicity)

## Architecture

```
microservices-design-patterns/
│── api-gateway/
│   ├── api-gateway-service/   # API Gateway service
│   │   ├── main.go            # API Gateway logic
│   │   ├── grpc_client.go     # gRPC Client for Order Service
│   │   ├── reverse_proxy.go   # HTTP Reverse Proxy for User Service
│   │   ├── proto/
│   │   │   ├── order.proto    # gRPC Service Definition
│   │   ├── generated/         # Compiled gRPC code
│   │   │   ├── order.pb.go
│   │   │   ├── order_grpc.pb.go
│   │   ├── Dockerfile
│   │   ├── docker-compose.yml
│
│── user-service/              # User Service (HTTP REST API)
│   ├── main.go
│   ├── routes.go
│   ├── Dockerfile
│   ├── docker-compose.yml
│
│── order-service/             # Order Service (gRPC API)
│   ├── main.go
│   ├── order_server.go
│   ├── proto/
│   │   ├── order.proto
│   ├── generated/
│   │   ├── order.pb.go
│   │   ├── order_grpc.pb.go
│   ├── Dockerfile
│   ├── docker-compose.yml
│
│── docker-compose.yml         # Global orchestration
```

## Setup & Installation

### 1. Clone the Repository

```
git clone https://github.com/your-repo/microservices-design-patterns.git
cd microservices-design-patterns/api-gateway
```

### 2️. Generate gRPC Code

```
cd api-gateway-service
protoc --go_out=generated --go-grpc_out=generated proto/order.proto
```

### 3️. Run Services with Docker Compose

```
docker-compose up --build
```

## API Gateway Functionality

🔹 User Service (HTTP Proxy)

Uses httputil.ReverseProxy to forward requests to user-service

Example: GET /api/user/profile → <http://user-service:8081/profile>

🔹 Order Service (gRPC Proxy)

Uses gRPC Client to call order-service

Passes JWT token via gRPC metadata

Example: GET /api/order/{orderID} → Calls order-service gRPC method

🔹 JWT Authentication

All routes require a valid JWT token

Token is extracted from Authorization header

Used in both HTTP & gRPC requests

### API Endpoints

User API (HTTP - Reverse Proxy)

#### Login (Get JWT Token)

```
TOKEN=$(curl -X POST "http://localhost:8080/api/user/login" -H "Content-Type: application/json" -d '{"username":"john","password":"1234"}' | jq -r .token)
```

#### Order API (gRPC via API Gateway)

Fetch Order Details (Requires JWT)

```
curl -X GET "http://localhost:8080/api/order/123" -H "Authorization: Bearer $TOKEN"
```
