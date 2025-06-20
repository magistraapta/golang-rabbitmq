package main

import (
	"log"
	"message-queue/payment-service/consumer"
)

func main() {
	err := consumer.ConsumerOrder()
	if err != nil {
		log.Fatalf("Failed to consume order: %v", err)
	}

	log.Println("Payment service started")
}
