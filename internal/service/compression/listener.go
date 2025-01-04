package compression

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"image-optimization-api/pkg/rabbitmq"
)

func NewListener(conn *amqp.Connection, service *Service) *QueueListener {
	return &QueueListener{
		conn:    conn,
		service: service,
	}
}

type QueueListener struct {
	conn    *amqp.Connection
	service *Service
}

func (s *QueueListener) ListenUpdates(ctx context.Context) error {
	msgs, err := rabbitmq.NewConsumer(s.conn)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			err = s.service.compressImage(ctx, msg.Body)
			if err != nil {
				zap.L().Error(fmt.Sprintf("Failed to compress images: %s", err.Error()))
			} else {
				zap.L().Info(fmt.Sprintf("Successfully compressed image"))
			}
		}
	}()

	go func() {
		<-ctx.Done()
		zap.L().Info("Shutting down queue listener...")
	}()

	return nil
}
