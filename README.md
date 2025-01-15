# Microservices Design Patterns in Go

## Overview
This repository demonstrates various microservices design patterns implemented using Go, the repository pattern, and containerized deployments. Each pattern is implemented as a separate microservice, allowing independent execution and scaling.

## Repository Structure
```
microservices-design-patterns/
│── api-gateway/
│   ├── src/
│   ├── Dockerfile
│   ├── docker-compose.yml
│   ├── README.md
│
│── event-processing/
│── service-registry/
│── saga/
│── container-orchestration/
│── circuit-breaker/
│
│── README.md
│── docker-compose.yml    # Global Orchestration
│── .gitignore
```

## Getting Started

### Prerequisites
- Docker and Docker Compose
- Go (>=1.23)
- Kubernetes (Minikube recommended)
- Kafka (for event-driven microservices)
- Prometheus & Grafana (for monitoring)

### Running Services
Each pattern can be run independently using Docker Compose. Navigate to the respective directory and execute:
```sh
docker compose up -d --build
```
To bring all services down:
```sh
docker compose down --volumes --remove-orphans
```

### Running Kubernetes Deployment
For a full Kubernetes deployment:
```sh
minikube start --memory=8192 --cpus=4
kubectl apply -f k8s/
```
To check the status:
```sh
kubectl get pods
```

## Available Microservices Design Patterns

### API Gateway
Handles request routing, composition, and authentication.

### Event Processing
Demonstrates asynchronous event-driven architecture using Kafka.

### Service Registry
Enables service discovery using tools like Consul or Zookeeper.

### Saga Pattern
Implements distributed transactions using orchestration or choreography.

### Container Orchestration
Demonstrates deployment using Kubernetes for scalability.

### Circuit Breaker
Prevents cascading failures in a distributed system.

## Monitoring & Observability
Prometheus and Grafana are used for monitoring microservices.

To access Grafana:
```sh
kubectl port-forward svc/grafana 3000:3000
```
Prometheus UI can be accessed at:
```sh
http://localhost:9091
```

## Contribution Guidelines
1. Fork the repository
2. Create a feature branch
3. Commit changes with descriptive messages
4. Open a pull request

## License
This repository is licensed under the MIT License.

