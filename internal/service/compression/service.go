package compression

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"image-optimization-api/internal/repository"
	"image-optimization-api/internal/service/image"
	"image-optimization-api/pkg/rabbitmq"
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
	msgs, err := rabbitmq.NewConsumer(s.conn)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			err = s.processMessage(ctx, msg.Body)
			if err != nil {
				zap.L().Error(fmt.Sprintf("Failed to process image: %s", err.Error()))
			} else {
				zap.L().Info(fmt.Sprintf("Successfully processed image: %s", msg.Body))
			}
		}
	}()

	go func() {
		<-ctx.Done()
		zap.L().Info("Shutting down consumer...")
	}()

	return nil
}

func (s *Service) processMessage(ctx context.Context, msg []byte) error {
	var obj image.UploadImageRequest

	err := obj.UnmarshalJSON(msg)
	if err != nil {
		return err
	}

	fmt.Println(obj)

	return nil
}
