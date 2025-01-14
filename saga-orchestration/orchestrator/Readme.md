# Microservices Design Patterns - Saga Pattern

## Overview

The Saga Pattern is used for managing distributed transactions across multiple microservices in an event-driven architecture. This implementation follows the Orchestration-Based Saga using Kafka as the message broker.

## Architecture

This repository includes the following microservices:

Saga Orchestrator: Manages the Saga workflow and ensures transaction consistency.

Order Service: Handles order creation and publishes events.

Payment Service: Handles payments and publishes success or failure events.

Kafka: Acts as the event broker to pass messages between services.

## Directory Structure

```
microservices-design-patterns/
│── saga/
│   ├── orchestrator/      # Saga Orchestrator
│   ├── order-service/     # Order Service
│   ├── payment-service/   # Payment Service
│   ├── Dockerfile         # Dockerfile for containerization
│── common/                # Shared configurations
│── tools/                 # CLI tools (if any)
│── tests/                 # Test scripts
│── docker-compose.yml     # Container orchestration
│── README.md              # Project documentation
```

## Setup and Deployment

1. Prerequisites

Ensure you have the following installed:

Docker

Docker Compose

Kafka & Zookeeper (handled via Docker)

2. Build and Start the Services

Run the following command to start all services:

docker-compose up --build -d

3. Verify Services Are Running

To check the status of running containers:

docker ps

4. Check Logs

To view logs for each service:

docker logs saga-orchestrator

docker logs order-service

docker logs payment-service

API Endpoints

Saga Orchestrator

Create Order

```
POST /create-order

Request Body:

{
  "order_id": "12345",
  "amount": 100.00
}
```

Example Request (cURL):

```
curl -X POST http://localhost:8080/create-order \
     -H "Content-Type: application/json" \
     -d '{"order_id": "12345", "amount": 100.00}'
```

## Kafka Topics & Events

Step | Service | Event

1|Order Service|OrderCreated

2|Saga Orchestrator|ProcessOrder

3|Payment Service|PaymentProcessed / PaymentFailed

4|Saga Orchestrator|OrderCompleted / OrderCancelled

## Debugging & Troubleshooting

1. Restart Services

```
docker-compose down && docker-compose up --build -d
```

2. Verify Kafka Topics

```
docker exec -it kafka bash
kafka-topics.sh --list --bootstrap-server kafka:9092
```

3. Ensure Services Can Connect to Kafka
```
docker exec -it saga-orchestrator sh
nc -z kafka 9092
```