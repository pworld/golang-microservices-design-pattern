# Microservices Design Patterns - Saga Synchronization

## Overview

The Saga Synchronization Pattern is used for managing synchronous distributed transactions across multiple microservices. Unlike Saga Orchestration, which is event-driven and asynchronous, Saga Synchronization ensures step-by-step coordination while maintaining data consistency.

## Architecture

This repository includes the following microservices:

Saga Synchronization Service: Ensures transactions execute synchronously with controlled compensations.

Saga Orchestrator: Manages the Saga workflow in an asynchronous manner.

Order Service: Handles order creation and interacts with synchronization logic.

Payment Service: Processes payments and ensures transactional consistency.

Kafka: Acts as the event broker for distributed services.

## Directory Structure

```
microservices-design-patterns/
│── saga/
│   ├── orchestrator/       # Saga Orchestrator Service
│   ├── synchronization/    # Saga Synchronization Service
│   ├── order-service/      # Order Service
│   ├── payment-service/    # Payment Service
│   ├── Dockerfile          # Dockerfile for containerization
│── common/                 # Shared configurations and utilities
│── tools/                  # CLI tools (if any)
│── tests/                  # Test scripts for saga flow
│── docker-compose.yml      # Container orchestration
│── README.md               # Project documentation
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

```
docker logs saga-synchronization

docker logs order-service

docker logs payment-service
```

API Endpoints

```
Saga Synchronization Service

Start Synchronous Transaction

POST /sync-transaction

Request Body:

{
  "order_id": "12345",
  "amount": 100.00
}
```

Example Request (cURL):
```
curl -X POST http://localhost:8083/sync-transaction \
     -H "Content-Type: application/json" \
     -d '{"order_id": "12345", "amount": 100.00}'
```

## Kafka Topics & Events

Step|Service|Event

1|Order Service|OrderCreated
2|Saga Synchronization|SyncTransactionStarted
3|Payment Service|PaymentProcessed / PaymentFailed
4|Saga Synchronization|SyncTransactionCompleted / SyncTransactionRollback

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
docker exec -it saga-synchronization sh
nc -z kafka 9092
```

4. Manually Create Kafka Topic if Missing

```
kafka-topics.sh --create --topic saga_synchronization --bootstrap-server kafka:9092 --partitions 1 --replication-factor 1
```