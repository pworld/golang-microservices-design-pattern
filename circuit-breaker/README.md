# Circuit Breaker Pattern in Microservices

## Overview

The Circuit Breaker Pattern prevents cascading failures in distributed systems by stopping calls to a failing service after a threshold of failures. Once the service recovers, the circuit breaker allows requests again. This project demonstrates the Circuit Breaker Pattern in a Golang microservices architecture using:

Gobreaker (Circuit Breaker Implementation)

Prometheus & Grafana (Monitoring & Alerting)

Jaeger (Distributed Tracing)

Docker Compose (Container Orchestration)

Directory Structure

```
microservices-design-patterns/
│── circuit-breaker-services/    # Circuit breaker service
│   ├── src/
│   │   ├── circuitbreaker/      # Gobreaker logic
│   │   │   ├── circuitbreaker.go
│   │   ├── handler/             # HTTP handlers
│   │   │   ├── handler.go
│   │   ├── service/             # Business logic + Prometheus metrics
│   │   │   ├── service.go
│   │   ├── main.go              # Entry point
│   ├── Dockerfile               # Docker container config
│   ├── docker-compose.yml       # Run independently
│
│── order-services/              # Order service (mock external API)
│   ├── src/
│   │   ├── handler/             # API handlers
│   │   │   ├── handler.go
│   │   ├── service/             # Business logic + Prometheus metrics
│   │   │   ├── service.go
│   │   ├── main.go              # Entry point
│   ├── Dockerfile               # Docker container config
│   ├── docker-compose.yml       # Run independently
│
│── grafana/                     # Grafana configuration
│   ├── provisioning/
│   │   ├── datasources/
│   │   │   ├── datasource.yml   # Auto-configures Prometheus
│
│── prometheus.yml               # Prometheus scraping config
│
│── docker-compose.yml            # Runs all services together
│── .gitignore                    # Ignore files
```

## Features

- Circuit Breaker Protection using Gobreaker
- Retry Mechanism for external API calls
- Prometheus Monitoring for tracking request failures
- Grafana Dashboards & Alerts when the circuit breaker triggers
- Jaeger Distributed Tracing for request tracking
- Fully Containerized with Docker Compose

## Setup & Installation

### 1. Clone the Repository

git clone https://github.com/your-repo/microservices-design-patterns.git
cd microservices-design-patterns

### 2. Build and Run Services

Run all services (Circuit Breaker, Order Service, Prometheus, Grafana, Jaeger):

docker-compose up --build

### 3. Access Services

Service

URL

Circuit Breaker API

http://localhost:8080/orders

Order Service API

http://localhost:8081/orders

Prometheus UI

http://localhost:9090

Grafana UI

http://localhost:3000


## How It Works

### Circuit Breaker Logic

API Calls: circuit-breaker-service calls order-service.

Failures Detected: If order-service fails multiple times, Gobreaker trips the circuit.

Block Requests: Further requests return an error without calling order-service.

Recovery: After a cooldown period, requests are allowed again.

### Monitoring & Alerts

Prometheus scrapes metrics (circuit_breaker_requests_total).

Grafana visualizes the data & sends alerts when failures exceed thresholds.

### Testing

1. Normal Requests

curl http://localhost:8080/orders

Response: Order data from order-service.

2. Simulate Failures

Stop the order-service to trigger the Circuit Breaker:

docker stop $(docker ps -q --filter "name=order-service")

Then, make repeated requests:

curl http://localhost:8080/orders

Response: 503 Service Unavailable (Circuit Breaker is tripped).

3. View Metrics in Prometheus

Go to http://localhost:9090 and run:

circuit_breaker_requests_total

4. View Alerts in Grafana

Open Grafana: http://localhost:3000

Check Alert Rules for Circuit Breaker Failures.

5. View Traces in Jaeger

Open Jaeger UI: http://localhost:16686

Click Find Traces to see request flow.