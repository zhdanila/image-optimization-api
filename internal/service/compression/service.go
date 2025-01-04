package compression

import (
	"context"
	"image-optimization-api/internal/repository"
)

func NewService(imageRepo *repository.Image) *Service {
	return &Service{
		imageRepo: imageRepo,
	}
}

type Service struct {
	imageRepo *repository.Image
}

func (s *Service) compressImage(ctx context.Context, msg []byte) error {
	var err error

	op := newOperationCompressImages(s, msg)
	if err = op.unmarshalBody(ctx); err != nil {
		return err
	}
	if err = op.compressImage(ctx); err != nil {
		return err
	}
	if err = op.uploadImages(ctx); err != nil {
		return err
	}

	return nil
}
