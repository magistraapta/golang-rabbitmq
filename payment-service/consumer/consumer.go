package consumer

import (
	"encoding/json"
	"log"
	"message-queue/payment-service/config"
	"message-queue/payment-service/models"
)

func ConsumerOrder() error {
	conn, ch, err := config.SetupRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"orders",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			var order models.Order
			if err := json.Unmarshal(msg.Body, &order); err != nil {
				log.Printf("Failed to unmarshal order: %v", err)
				continue
			}
			log.Printf("processing order %s with amount %d", order.ID, order.Amount)

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")

	<-forever

	select {}
}
