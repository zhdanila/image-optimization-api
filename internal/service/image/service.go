package image

import (
	"context"
	"image-optimization-api/internal/repository"
)

func NewService(
	imageRepo *repository.Image,
) *Service {
	return &Service{
		imageRepo: imageRepo,
	}
}

type Service struct {
	imageRepo *repository.Image
}

func (s *Service) UploadImage(ctx context.Context, obj *UploadImageRequest) (*UploadImageResponse, error) {
	var err error

	op := newOperationUploadImage(s, obj)

	if err = op.uploadImage(ctx); err != nil {
		return nil, err
	}

	return op.respond(), nil
}
