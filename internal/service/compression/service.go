package compression

import (
	"context"
	"fmt"
	"image-optimization-api/internal/repository"
	"image-optimization-api/internal/service/image"
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
	var obj image.UploadImageRequest

	err := obj.UnmarshalJSON(msg)
	if err != nil {
		return err
	}

	fmt.Println(obj)

	return nil
}
