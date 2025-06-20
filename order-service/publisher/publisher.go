package publisher

import (
	"encoding/json"
	"message-queue/order-service/config"
	"message-queue/order-service/models"

	"github.com/rabbitmq/amqp091-go"
)

func PublisherOrder(order models.Order) error {
	conn, ch, err := config.SetupRabbitMQ()

	if err != nil {
		return err
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

	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
