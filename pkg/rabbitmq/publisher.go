package rabbitmq

import (
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

const ImageQueueName = "image_upload_queue"

func PublishToQueue(conn *amqp.Connection, body []byte) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		ImageQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}

	zap.L().Info("Message sent to queue")

	return nil
}
