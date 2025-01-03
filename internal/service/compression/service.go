package compression

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"image-optimization-api/internal/repository"
)

func NewService(
	conn *amqp.Connection,
	imageRepo *repository.Image,
) *Service {
	return &Service{
		conn:      conn,
		imageRepo: imageRepo,
	}
}

type Service struct {
	conn      *amqp.Connection
	imageRepo *repository.Image
}

func (s *Service) ListenUpdates(ctx context.Context) error {
	ch, err := s.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"image_optimization_queue", // Queue name
		true,                       // Durable
		false,                      // Auto-delete
		false,                      // Exclusive
		false,                      // No-wait
		nil,                        // Arguments
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		queue.Name, // Queue name
		"",         // Consumer name
		true,       // Auto-ack
		false,      // Exclusive
		false,      // No-local
		false,      // No-wait
		nil,        // Arguments
	)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case msg := <-msgs:
			err = s.processMessage(ctx)
			if err != nil {
				zap.L().Error(fmt.Sprintf("Failed to process image: %s", err.Error()))
			} else {
				zap.L().Error(fmt.Sprintf("Successfully processed image: %s", msg.Body))
			}
		}
	}
}

func (s *Service) processMessage(ctx context.Context) error {

}
