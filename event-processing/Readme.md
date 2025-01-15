# Event Processing Microservices with Kafka in Go

## Overview
This project implements an **event-driven microservices architecture** using **Kafka** for messaging, **gRPC/REST APIs** for monitoring, and **Kubernetes** for deployment. The system consists of:

- **Producer Service**: Sends events to Kafka.
- **Consumer Service**: Listens to Kafka topics and processes events.
- **Kafka & Zookeeper**: Handles message brokering and coordination.
- **Prometheus & Grafana**: Enables observability and monitoring.
- **Docker & Kubernetes**: Containerized deployment and orchestration.

---

## Directory Structure
```
-  golang-microservices-design-pattern
└── -  event-processing
    ├── -  consumer-service
    │   ├── Dockerfile
    │   ├── go.mod
    │   ├── go.sum
    │   ├── main.go
    │   ├── prometheus_metrics.go
    ├── -  producer-service
    │   ├── Dockerfile
    │   ├── go.mod
    │   ├── go.sum
    │   ├── main.go
    │   ├── prometheus_metrics.go
    ├── -  monitoring
    │   ├── prometheus.yml
    │   ├── grafana-config.json
    ├── -  k8s
    │   ├── consumer-deployment.yaml
    │   ├── producer-deployment.yaml
    │   ├── kafka-deployment.yaml
    │   ├── zookeeper-deployment.yaml
    │   ├── grafana-deployment.yaml
    │   ├── prometheus-deployment.yaml
    ├── docker-compose.yml
    ├── README.md
```

---

## Setup and Installation

### **1️⃣ Prerequisites**
Ensure you have the following installed:
- **Docker** & **Docker Compose**
- **Go (>=1.23)**
- **Minikube** (for Kubernetes testing)
- **kubectl**

### **2️⃣ Running with Docker Compose**
```sh
docker compose up -d --build
```
Check running services:
```sh
docker compose ps
```

To shut down:
```sh
docker compose down --volumes --remove-orphans
```

### **3️⃣ Running in Kubernetes**
Start Minikube:
```sh
minikube start --memory=8192 --cpus=4
```
Deploy all services:
```sh
kubectl apply -f k8s/
```
Check the status of pods:
```sh
kubectl get pods
```
Port-forward **Grafana** for monitoring:
```sh
kubectl port-forward svc/grafana 3000:3000
```

---

## Services

### **1️. Producer Service**
- Sends events to Kafka.
- API Endpoint:
```sh
curl -X POST http://localhost:8080/event -H "Content-Type: application/json" -d '{"message": "Hello Kafka!"}'
```
- Logs:
```sh
docker logs producer-service --tail=50
```

### **2️. Consumer Service**
- Listens to Kafka and processes messages.
- Logs:
```sh
docker logs consumer-service --tail=50
```

### **3️. Kafka & Zookeeper**
- View topics:
```sh
docker exec -it kafka kafka-topics --list --bootstrap-server localhost:9092
```
- Describe a topic:
```sh
docker exec -it kafka kafka-topics --describe --topic event-topic --bootstrap-server localhost:9092
```

### **4️. Monitoring: Prometheus & Grafana**
- Prometheus UI: [http://localhost:9091](http://localhost:9091)
- Grafana UI: [http://localhost:3000](http://localhost:3000)
  - Default login: `admin / admin`

---

## Troubleshooting

### **1️. Verify Container Logs**
```sh
docker compose logs --tail=50
```

### **2️. Debug Inside Containers**
```sh
docker exec -it consumer-service /bin/sh
```

### **3️. Rebuild with No Cache**
```sh
docker compose build --no-cache
```

### **4️. Check Kubernetes Pods**
```sh
kubectl get pods -A
kubectl logs -l app=consumer --tail=50
```

### **5️. Kafka Debugging**
- List topics:
```sh
kubectl exec -it kafka -- kafka-topics --list --bootstrap-server kafka:9092
```
- Consume messages:
```sh
kubectl exec -it kafka -- kafka-console-consumer --topic event-topic --from-beginning --bootstrap-server kafka:9092
```