# Message Queue System

A microservices-based message queue system built with Go and RabbitMQ, demonstrating asynchronous communication between services.

check the article associated with this project [here](https://medium.com/@magistraagis/implement-message-queue-using-rabbitmq-2902d1f0014d)

## Project Overview

This project implements a distributed system with two microservices:
- **Order Service**: Creates orders and publishes them to a message queue
- **Payment Service**: Consumes orders from the queue and processes payments

The system uses RabbitMQ as the message broker to enable asynchronous, decoupled communication between services.

## Architecture

```
┌─────────────────┐    HTTP POST    ┌─────────────────┐
│   Client        │ ──────────────► │  Order Service  │
│                 │                 │   (Port 8080)   │
└─────────────────┘                 └─────────────────┘
                                              │
                                              │ Publishes to Queue
                                              ▼
                                    ┌─────────────────┐
                                    │   RabbitMQ      │
                                    │   (Port 5672)   │
                                    │   Queue: orders │
                                    └─────────────────┘
                                              │
                                              │ Consumes from Queue
                                              ▼
                                    ┌─────────────────┐
                                    │ Payment Service │
                                    │   (Consumer)    │
                                    └─────────────────┘
```

## Project Structure

```
message-queue/
├── go.mod                          # Go module dependencies
├── go.sum                          # Dependency checksums
├── order-service/                  # Order creation microservice
│   ├── main.go                     # Service entry point
│   ├── config/
│   │   └── rabbit_mq.go           # RabbitMQ connection setup
│   ├── handler/
│   │   └── handler.go             # HTTP request handlers
│   ├── models/
│   │   └── order.go               # Order data structure
│   ├── publisher/
│   │   └── publisher.go           # Message publishing logic
│   └── router/
│       └── router.go              # HTTP routing configuration
├── payment-service/                # Payment processing microservice
│   ├── main.go                     # Service entry point
│   ├── config/
│   │   └── rabbitmq.go            # RabbitMQ connection setup
│   ├── consumer/
│   │   └── consumer.go            # Message consumption logic
│   └── models/
│       ├── order.go               # Order data structure
│       └── payment.go             # Payment data structure
└── readme.md                       # This file
```

## Services Description

### Order Service (`order-service/`)

**Purpose**: Handles order creation and publishes orders to RabbitMQ queue.

**Components**:
- **`main.go`**: Starts the HTTP server on port 8080
- **`router/router.go`**: Sets up HTTP routes (POST `/orders`)
- **`handler/handler.go`**: Handles order creation requests, generates UUIDs, and triggers publishing
- **`publisher/publisher.go`**: Publishes order messages to RabbitMQ queue
- **`config/rabbit_mq.go`**: Manages RabbitMQ connection and channel setup
- **`models/order.go`**: Defines the Order struct with ID and Amount fields

**API Endpoint**:
- `POST /orders` - Creates a new order and publishes it to the queue

### Payment Service (`payment-service/`)

**Purpose**: Consumes orders from RabbitMQ queue and processes payments.

**Components**:
- **`main.go`**: Starts the consumer service
- **`consumer/consumer.go`**: Consumes messages from the queue and processes orders
- **`config/rabbitmq.go`**: Manages RabbitMQ connection and channel setup
- **`models/order.go`**: Order data structure for deserialization
- **`models/payment.go`**: Payment data structure

## Data Models

### Order Model
```go
type Order struct {
    ID     string `json:"id"`
    Amount int    `json:"amount"`
}
```

### Payment Model
```go
type Payment struct {
    OrderID string `json:"order_id"`
    Amount  int    `json:"amount"`
}
```

## Dependencies

- **`github.com/google/uuid`**: UUID generation for order IDs
- **`github.com/gorilla/mux`**: HTTP routing and middleware
- **`github.com/rabbitmq/amqp091-go`**: RabbitMQ client library

## Prerequisites

1. **Go 1.24.0 or higher**
2. **RabbitMQ Server** running on `localhost:5672`
   - Default credentials: `guest:guest`

## Getting Started

### 1. Install Dependencies
```bash
go mod download
```

### 2. Start RabbitMQ
Make sure RabbitMQ is running on your system:
```bash
# On macOS with Homebrew
brew services start rabbitmq

# On Ubuntu/Debian
sudo systemctl start rabbitmq-server

# Or using Docker
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management
```

### 3. Start the Services

**Terminal 1 - Start Payment Service (Consumer)**:
```bash
cd payment-service
go run main.go
```

**Terminal 2 - Start Order Service**:
```bash
cd order-service
go run main.go
```

### 4. Test the System

Create an order using curl:
```bash
curl -X POST http://localhost:8080/orders \
  -H "Content-Type: application/json" \
  -d '{"amount": 100}'
```

Expected response:
```json
{"message": "Order created successfully"}
```

The payment service will log the order processing:
```
processing order <uuid> with amount 100
```

## Message Flow

1. **Client** sends POST request to Order Service with order details
2. **Order Service** generates a unique ID and publishes the order to RabbitMQ queue
3. **Payment Service** consumes the order from the queue and processes it
4. **Payment Service** logs the order processing details

## Queue Configuration

- **Queue Name**: `orders`
- **Durability**: Non-durable (messages lost on restart)
- **Auto-delete**: False
- **Exclusive**: False
- **Auto-acknowledge**: True

## Error Handling

- Connection failures are logged and handled gracefully
- JSON unmarshaling errors are logged but don't stop the consumer
- HTTP errors return appropriate status codes

## Development Notes

- Both services use the same RabbitMQ connection configuration
- The system is designed for demonstration purposes
- In production, consider adding:
  - Message persistence
  - Error queues
  - Retry mechanisms
  - Health checks
  - Metrics and monitoring
  - Security configurations

## Troubleshooting

1. **Connection refused**: Ensure RabbitMQ is running on port 5672
2. **Queue not found**: The queue is created automatically when first accessed
3. **Service won't start**: Check Go version and dependencies with `go mod tidy`
