package image

import (
	"context"
	"github.com/streadway/amqp"
	"image-optimization-api/internal/repository"
	"image-optimization-api/pkg/server"
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

func (s *Service) UploadImage(ctx context.Context, obj *UploadImageRequest) (*server.EmptyResponse, error) {
	var err error

	op := newOperationQueuePublish(s, obj)

	if err = op.queuePublish(ctx); err != nil {
		return nil, err
	}

	return op.respond(), nil
}

func (s *Service) GetImage(ctx context.Context, obj *GetImageRequest) (*GetImageResponse, error) {
	var err error

	op := newOperationGetImage(s, obj)

	if err = op.prepareImageID(ctx); err != nil {
		return nil, err
	}
	if err = op.getImage(ctx); err != nil {
		return nil, err
	}

	return op.respond(), nil
}

func (s *Service) ListImages(ctx context.Context, obj *ListImageRequest) (*ListImageResponse, error) {
	var err error

	op := newOperationListImages(s, obj)

	if err = op.listImages(ctx); err != nil {
		return nil, err
	}

	return op.respond(), nil
}
