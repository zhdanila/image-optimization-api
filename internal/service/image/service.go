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

func (s *Service) GetImage(ctx context.Context, obj *GetImageRequest) (*GetImageResponse, error) {
	var err error

	op := newOperationGetImage(s, obj)

	if err = op.getImage(ctx); err != nil {
		return nil, err
	}

	return op.respond(), nil
}
